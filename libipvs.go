package libipvs

/*
#include "libipvs.h"
*/
import "C"

/* init socket and get ipvs info */
func ipvs_init() int {
	return int(C.ipvs_init())
}

/* get ipvs info separately */
func ipvs_getinfo() int {
	return int(C.ipvs_getinfo())
}

/* get the version number */
func ipvs_version() uint {
	return uint(C.ipvs_version())
}

/* flush all the rules */
func ipvs_flush() int {
	return int(C.ipvs_flush())
}

/* add a virtual service */
func ipvs_add_service(svc *C.ipvs_service_t) int {
	return int(C.ipvs_add_service(svc))
}

/* update a virtual service with new options */
func ipvs_update_service(svc *C.ipvs_service_t) int {
	return int(C.ipvs_update_service(svc))
}

/* delete a virtual service */
func ipvs_del_service(svc *C.ipvs_service_t) int {
	return int(C.ipvs_del_service(svc))
}

/* zero the counters of a service or all */
func ipvs_zero_service(svc *C.ipvs_service_t) int {
	return int(C.ipvs_zero_service(svc))
}

/* add a destination server into a service */
func ipvs_add_dest(svc *C.ipvs_service_t, dest *C.ipvs_dest_t) int {
	return int(C.ipvs_add_dest(svc, dest))
}

/* update a destination server with new options */
func ipvs_update_dest(svc *C.ipvs_service_t, dest *C.ipvs_dest_t) int {
	return int(C.ipvs_update_dest(svc, dest))
}

/* remove a destination server from a service */
func ipvs_del_dest(svc *C.ipvs_service_t, dest *C.ipvs_dest_t) int {
	return int(C.ipvs_del_dest(svc, dest))
}

/* set timeout */
func ipvs_set_timeout(to *C.ipvs_timeout_t) int {
	return int(C.ipvs_set_timeout(to))
}

/* start a connection synchronizaiton daemon (master/backup) */
func ipvs_start_daemon(dm *C.ipvs_daemon_t) int {
	return int(C.ipvs_start_daemon(dm))
}

/* stop a connection synchronizaiton daemon (master/backup) */
func ipvs_stop_daemon(dm *C.ipvs_daemon_t) int {
	return int(C.ipvs_stop_daemon(dm))
}
