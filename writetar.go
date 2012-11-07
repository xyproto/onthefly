package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/hoisie/web"
)

func writeFile(w *tar.Writer, pkgdir, filename string, contents map[string]string) {
	contentString := ""
	for key, value := range contents {
		contentString += "%" + strings.ToUpper(key) + "%\n" + value + "\n\n"
	}

	data := bytes.NewBufferString(contentString)
	size := int64(len(contentString))

	// Create the header
	hdr := new(tar.Header)

	hdr.Name = pkgdir + "/" + filename
	hdr.Mode = 0666
	//hdr.Uid = 1000
	//hdr.Gid = 1000
	hdr.Size = size
	hdr.ModTime = time.Now()
	//hdr.Typeflag = 0 // ??
	//hdr.Linkname = "" // ??
	//hdr.Uname = "alexander"
	//hdr.Gname = "users"
	//hdr.Devmajor = 0 // ??
	//hdr.Devminor = 0 // ??
	//hdr.AccessTime = ? 
	//hdr.ChangeTime = ?

	// Write the header to the tar file
	if err := w.WriteHeader(hdr); err != nil {
		log.Fatal("error when writing header for " + pkgdir)
	}

	// Write the data to the tar file
	io.Copy(w, data)
}

func writeTar(fw io.Writer) {
	tw := tar.NewWriter(fw)

	for _, pkgdir := range []string{"zlib-1.2.7-2"} {
		fmt.Println(pkgdir)

		depends := make(map[string]string)
		depends["depends"] = "glibc"
		writeFile(tw, pkgdir, "depends", depends)

		desc := make(map[string]string)
		desc["filename"] = pkgdir + "-x86_64.pkg.tar.xz"
		desc["name"] = "zlib"
		desc["version"] = "1.2.7-2"
		desc["csize"] = "79156"
		desc["isize"] = "368640"
		desc["md5sum"] = "986040c038621e82e327480c4f28c804"
		desc["sha256sum"] = "a3d074554d0cef230c816002e6be995a26a927d2716126af63fd3a60637bc164"
		desc["pgpsig"] = "iQEcBAABAgAGBQJPokOtAAoJEH8tQ0uXQeisqjMH/207fWE+1vcpz5xQiOLq2GmTwILUki7uAL9RQ1AeD2KO7b8KLpy6fx881/J3LzwoNLP5ygcKyjaYBGZ7IG4E90Wb+O/Emiz3MyXya+ZuorQrg18LcKoy5SkK+z823LhmzN0eXuoE1l2b/b2R7L5VJ97Lx0G0YcNenD6mYKwYA4tG3XG70kAjyGo4Aunq6ipaIfOJIRs3n5MRZFOSb7LFsSsonL7GAnGkFrHFRPIBNGk6K+wl0frOlsv6qdMuWsHqC93Y1C0i0ZCAIxJiPbbr+odCU1nFapRZKJPaQSq/m1vxuXUK3VAYkC5nciQ+VPSlllucnMbPZZZh2k67RViz6gY="
		desc["url"] = "http://www.zlib.net/"
		desc["license"] = "custom"
		desc["arch"] = "x86_64"
		desc["builddate"] = "1336034033"
		desc["packager"] = "Pierre Schmitz <pierre@archlinux.de>"
		writeFile(tw, pkgdir, "desc", desc)
	}

	tw.Close()
}

func writeTarGzToFile(tarfilename string) error {
	var perm os.FileMode = 0666
	f, err := os.OpenFile(tarfilename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	gz := gzip.NewWriter(f)
	writeTar(gz)
	gz.Close()
	f.Close()
	return nil
}

func writeTarGz(rw web.ResponseWriter) {
	gz := gzip.NewWriter(rw)
	writeTar(gz)
	//gz.Close()
}

func test_writetar() {
	writeTarGzToFile("test.tar.gz")
}
