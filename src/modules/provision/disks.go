package provision

import "vmctl/src/config"

func initDisks(config ProvisionOptions, cfg *config.AppConfig) {
	// # if the env is a dev one, then we don't want to use a real disk
	// # and we want to tag all state dirs with 'dev'
	// if [[ "$DEVELOPMENT" == false ]]; then
	// 	set_thinpool="${DEFAULT_THINPOOL:=$thinpool}"
	// 	do_all_direct_lvm "$disk" "$set_thinpool"
	// else
	// 	set_thinpool="${DEFAULT_DEV_THINPOOL:=$thinpool}"
	// 	do_all_devpool "$set_thinpool"
	// fi
}
