
## DIRECT_LVM
#
#
do_all_direct_lvm() {
	local disk="$1"
	local thinpool="$2"

	say "Setting up direct_lvm thinpool $thinpool"

	if [[ -z "$disk" ]]; then
		say_warn "WARNING: -d/--disk has not been set. If you continue, the" \
			"script will attempt to detect a free disk for formatting. Any data" \
			"will be lost."
		get_user_confirmation "Are you sure you wish to continue? (y/n) " || die "Aborted."

		disk=$(find_free_disk || die "Could not detect free disk")
	fi

	disk_name=$(basename "$disk")
	local disk_path="/dev/$disk_name"

	say "Will use $disk_path for direct-lvm thinpool $thinpool"
	say_warn "All existing data on $disk_path will be overwritten."
	get_user_confirmation || die "Aborted."

	create_physical_volume "$disk_path"
	create_volume_group "$disk_path" "$thinpool"
	create_logical_volume "$thinpool"
	apply_lvm_profile "$thinpool"
	monitor_lvm_profile "$thinpool" || die "failed to monitor lvm profile"

	say "Thinpool $thinpool-thinpool is ready for use"
}

# Naively find a spare block device which is not in use
# This is really unsafe as it only looks at a couple of things to decide anything
# It is a much better idea to use the --disk flag and pass in something you
# know will be available
find_free_disk() {
	disks=("$(lsblk -o NAME,TYPE | awk '/disk/ {print $1}')")

	# shellcheck disable=SC2068
	for d in ${disks[@]}; do
		if ! is_mounted "$d" && ! is_partitioned "$d"; then
			echo "$d" && return 0
		fi
	done

	return 1
}

# Check whether given device is mounted
is_mounted() {
	local device_name="$1"
	findmnt -rno TARGET "/dev/$device_name" >/dev/null
}

# Check whether given device is partitioned
is_partitioned() {
	local device_name="$1"
	sfdisk -d "/dev/$device_name" &>/dev/null
}

# Create a physical volume on the given device
create_physical_volume() {
	local disk_path="$1"

	# if already exists, do nothing
	if [[ $(pvdisplay 2>/dev/null) != *"$disk_path"* ]]; then
		pvcreate -q "$disk_path" || die "failed to create physical volume on $disk_path"
		say "Created physical volume on $disk_path"
		return 0
	fi

	say "Physical volume on $disk_path already exists"
}

# Create a volume group on the given device for the thinpool
create_volume_group() {
	local disk_path="$1"
	local thinpool="$2"

	# if already exists, do nothing
	if [[ $(vgdisplay 2>/dev/null) != *"$thinpool"* ]]; then
		vgcreate -q "$thinpool" "$disk_path" || die "failed to create volume group on $disk_path"
		say "Created volume group on $disk_path"
		return 0
	fi

	say "Volume group on $disk_path already exists"
}

# Format the volume for thinpool storage
create_logical_volume() {
	local volume_group="$1"

	# if already exists, do nothing
	if [[ $(lvdisplay 2>/dev/null) != *"$volume_group"* ]]; then
		lvcreate -q --wipesignatures y -n thinpool "$volume_group" -l 95%VG || die "Failed to create logical volume for thinpool data"
		say "Created logical volume for $volume_group thinpool data"

		lvcreate -q --wipesignatures y -n thinpoolmeta "$thinpool" -l 1%VG || die "Failed to create logical volume for thinpool metadata"
		say "Created logical volume for $volume_group thinpool metadata"

		lvconvert -q -y \
			--zero n \
			-c 512K \
			--thinpool "$volume_group"/thinpool \
			--poolmetadata "$volume_group"/thinpoolmeta || die "Failed to convert logical volumes to thinpool storage"
		say "Converted logical volumes to thinpool storage"

		return 0
	fi

	say "Logical volume for $volume_group thinpool already exists"
}

# Create and apply the lvm profile for the thinpool
apply_lvm_profile() {
	local thinpool="$1"
	local profile="$THINPOOL_PROFILE_PATH/$thinpool-thinpool.profile"

	if [[ ! -f "$profile" ]]; then
		cat <<'EOF' >>"$profile"
activation {
thin_pool_autoextend_threshold=80
thin_pool_autoextend_percent=20
}
EOF
		say "Written lvm profile to $profile"
	fi

	# if already exists, do nothing
	if [[ $(lvs 2>/dev/null) != *"$thinpool"* ]]; then
		lvchange -q --metadataprofile "$thinpool-thinpool" "$thinpool"/thinpool || die "Could not apply lvm profile $profile"
		say "Applied lvm profile for $thinpool-thinpool"
		return 0
	fi

	say "LVM profile for $thinpool-thinpool already applied"
}

# Try 5 times to ensure the lvm profile is monitored
monitor_lvm_profile() {
	local thinpool="$1"

	for _ in $(seq 5); do
		if [[ $(lvs -o+seg_monitor 2>/dev/null) != *"not monitored"* ]]; then
			say "Successfully monitored ${thinpool}-thinpool profile"
			return
		fi
		lvchange --monitor y "${thinpool}/thinpool"
	done

	die -c 1
}
