package sqlite3

/*
#cgo linux,amd64 CFLAGS: -msse4.1
#cgo CFLAGS: -DUSE_LIBSQLITE3 -DCODEC_TYPE=CODEC_TYPE_AES128 -DSQLITE_HAS_CODEC=1
*/
import "C"
