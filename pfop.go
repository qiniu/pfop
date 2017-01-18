package main

import (
	"flag"
	"fmt"
	"github.com/qiniu/api.v6/auth/digest"
	"github.com/qiniu/rpc"
	"net/http"
)

const (
	PFOP_NB_URL   = "http://api.qiniu.com/pfop/"
	PREFOP_NB_URL = "http://api.qiniu.com/status/get/prefop?id=%s"
)

const (
	PFOP_AWS_URL   = "http://api.gdipper.com/pfop/"
	PREFOP_AWS_URL = "http://api.gdipper.com/status/get/prefop?id=%s"
)

const (
	PFOP_BC_URL   = "http://api-z1.qiniu.com/pfop/"
	PREFOP_BC_URL = "http://api-z1.qiniu.com/status/get/prefop?id=%s"
)

const (
	PFOP_HN_URL   = "http://api-z2.qiniu.com/pfop/"
	PREFOP_HN_URL = "http://api-z2.qiniu.com/status/get/prefop?id=%s"
)

const (
	PFOP_NA0_URL   = "http://api-na0.qiniu.com/pfop/"
	PREFOP_NA0_URL = "http://api-na0.qiniu.com/status/get/prefop?id=%s"
)

type PfopResult struct {
	PersistentId string `json:"persistentId,omitempty"`
}

func pfop(ak, sk, bucket, key, fops string, notifyUrl, pipeline, zone string, force bool) {
	mac := digest.Mac{
		ak,
		[]byte(sk),
	}

	t := digest.NewTransport(&mac, nil)
	client := &http.Client{Transport: t}
	rpcClient := rpc.Client{client}
	pfopResult := PfopResult{}

	pfopParams := map[string][]string{
		"bucket": []string{bucket},
		"key":    []string{key},
		"fops":   []string{fops},
	}
	if notifyUrl != "" {
		pfopParams["notifyURL"] = []string{notifyUrl}
	}
	if pipeline != "" {
		pfopParams["pipeline"] = []string{pipeline}
	}
	if force {
		pfopParams["force"] = []string{"1"}
	}

	pfopUrl := PFOP_NB_URL
	prefopUrl := PREFOP_NB_URL

	switch zone {
	case "aws":
		pfopUrl = PFOP_AWS_URL
		prefopUrl = PREFOP_AWS_URL
	case "bc":
		pfopUrl = PFOP_BC_URL
		prefopUrl = PREFOP_BC_URL
	case "hn":
		pfopUrl = PFOP_HN_URL
		prefopUrl = PREFOP_HN_URL
	case "na0":
		pfopUrl = PFOP_NA0_URL
		prefopUrl = PREFOP_NA0_URL
	default:
		pfopUrl = PFOP_NB_URL
		prefopUrl = PREFOP_NB_URL
	}

	err := rpcClient.CallWithForm(nil, &pfopResult, pfopUrl, pfopParams)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fmt.Sprintf(fmt.Sprintf("See %s", prefopUrl), pfopResult.PersistentId))
	}
}

func main() {
	var ak string
	var sk string
	var bucket string
	var key string
	var fops string
	var notifyUrl string
	var pipeline string
	var force bool
	var zone string

	flag.StringVar(&ak, "ak", "", "access key")
	flag.StringVar(&sk, "sk", "", "secret key")
	flag.StringVar(&bucket, "bucket", "", "bucket name")
	flag.StringVar(&key, "key", "", "file name")
	flag.StringVar(&fops, "fops", "", "joined fop")
	flag.StringVar(&notifyUrl, "url", "", "notify url")
	flag.StringVar(&pipeline, "pipe", "", "pipeline to use")
	flag.StringVar(&zone, "zone", "nb", "api zone [nb, bc, hn or na0]")
	flag.BoolVar(&force, "force", false, "force to redo")

	flag.Usage = func() {
		fmt.Println(`Usage of pfop:
  -ak="": access key
  -sk="": secret key
  -bucket="": bucket name
  -key="": file key
  -fops="": joined fops
  -pipe="": pipeline to use
  -url="": notify url
  -zone="nb": api zone [nb, bc, hn, na0]
  -force: force to redo
`)
	}
	flag.Parse()

	if ak == "" {
		fmt.Println("Invalid pfop parameter 'ak'")
		return
	}

	if sk == "" {
		fmt.Println("Invalid pfop parameter 'sk'")
		return
	}

	if bucket == "" {
		fmt.Println("Invalid pfop parameter 'bucket'")
		return
	}

	if key == "" {
		fmt.Println("Invalid pfop parameter 'key'")
		return
	}

	if fops == "" {
		fmt.Println("Invalid pfop parameter 'fops'")
		return
	}

	if !(zone == "nb" || zone == "bc" || zone == "hn" || zone == "na0" || zone == "aws") {
		fmt.Println("Invalid pfop parameter 'zone'")
		return
	}

	if pipeline == "" {
		fmt.Println("Tip: specify pfop parameter 'pipe' to ensure processing speed!")
	}

	pfop(ak, sk, bucket, key, fops, notifyUrl, pipeline, zone, force)
}
