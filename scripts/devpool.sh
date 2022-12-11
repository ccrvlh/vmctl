
## DEVPOOL
#
#
do_all_devpool() {
	local thinpool="$1-thinpool"

	say "Will create loop-back thinpool $thinpool"

	create_sparse_file "$DEVPOOL_DATA" "$DATA_SPARSE_SIZE"
	create_sparse_file "$DEVPOOL_METADATA" "$METADATA_SPARSE_SIZE"

	say "Associating loop devices with sparse files"
	datadev=$(associate_loop_device "$DEVPOOL_DATA")
	metadev=$(associate_loop_device "$DEVPOOL_METADATA")
	say "Loop devices $datadev and $metadev associated"

	create_dev_thinpool "$thinpool" "$datadev" "$metadev"

	say "Dev thinpool creation complete"
}

# Create the a sparse file which will be used to back a loop device
create_sparse_file() {
	local file="$1"
	local size="$2"

	say "Creating sparse file $file of size $size"
	if [[ ! -f "$file" ]]; then
		touch "$file"
		truncate -s "$size" "$file" || die "Failed to create sparse file $file"
	fi

	say "Sparse file $file created"
}

# Assign a loop device to the given sparse file
associate_loop_device() {
	local sparse_file="$1"

	device="$(losetup --output NAME --noheadings --associated "$sparse_file")"
	if [[ -z "$device" ]]; then
		device=$(losetup --find --show "$sparse_file" || die "Failed to associate loop device with $sparse_file")
	fi

	echo "$device"
}

# Create the thinpool with the loop devices if it does not already exist
create_dev_thinpool() {
	local thinpool="$1"
	local datadev="$2"
	local metadev="$3"

	say "Creating thinpool $thinpool with devices $datadev and $metadev"

	datasize="$(blockdev --getsize64 -q "$datadev")"
	length_sectors=$(bc <<<"$datasize/$SECTORSIZE")
	thinp_table="0 $length_sectors thin-pool $metadev $datadev $DATA_BLOCK_SIZE $LOW_WATER_MARK 1 skip_block_zeroing"

	if ! dmsetup reload "$thinpool" --table "$thinp_table" 2>/dev/null; then
		dmsetup create "$thinpool" --table "$thinp_table" || die "failed to create dev thinpool $thinpool"
	fi

	say "Thinpool $thinpool created"
}
