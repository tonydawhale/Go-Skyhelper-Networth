package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/beito123/nbt"
)

func main() {
	var base64String = "H4sIAAAAAAAAANWVzW7bRhSFR3aa2KpttSiyyabjpEYXMW1S/NdOVmQkSCwXqZ2gK2PIuZQGImcEcuREz9GH0HvowYreoQQ7NlS0KGIU5YIiR4d3Ls/5cNkkZJs0RJMQYm2QDcEbvzfINz01lbrRJJuaDbfJI5DpiJhjk2y/FhxOczas8PaPJnnCRTXJ2QxV71QJW7i6Q75fzOOL0VRyKHNVcvrhDXmxmIf9ayhndDFPXTyFI6ErqiRltFCy0lAekue4XOlSjAEFkIvhSEshh7h6SDmwHK/JPt4t5sy3D0wRlVE9Aoq1fq4oZwUbwhG28KIWhd3JJJ/Rnqp0B8Vu27Zp//OEvoNryE3/+yuNSFmSAzZjZPEZ5AD0I7CJknVLlxXgLmLZraRCQ0FFfdmV1yInLdRoRVm9m9D7WLmJdRbz/H33fX+LPBqwAmpT+mgkkxo4PVFqTJqk1f+sS9bV+NLJVEO1RbZVKYZCXrAheXI5eDs4/zjYMrmQvf6g97o7uOi/ujo5P3/bJDuwrFaA1NUm+VbfOo4dPCa48XSKD/7kcIgyiEMrjULH8qLIs2I/TSw3Bs9pR20vclPcWIsCKs2KCdnzj53jtkPdju3R7hkhG+Txq9rcmgLylUDZJU/Rk1NgGGFJT1luAjawGNPfyLQEVkFFR+oTBjwc0Zma0hRdz1BJfkRNAhmWopqNzYNmecUATWbkYAlKgD8YFacl8GmK5e7Jdpcy1zZ6w87ztewE99B5thYdTFVX/yNmWtnS/Ktsaf5dbrIgC0M3ta04DELLSxJuRUGG8EQJsDBwoH2Pm5Z/3PYNOH7HDteA85W4aRn30CqObp0wOabL8fIru8Z4F/PAuZkNSExJUyWkyaFGhpsXPqJdzoUWSuJrzw4pSCgEVDUyvFQTU+SlfeSvHsXVTyOQdCzyHPgtI+lvKyCVxBBHuD116GWOfjANhpC0f+v2FyDs/yVl/j+irFsWqvy3lO0s5hn61js/OzsffMFZ9pCc7U5XrlwlmBd20bilLPEd2wuCthU5NlieG4CVuBEOK8cJ7YQxcFh8dzoFy+kUd7yI/vJg06lFvruhrKeKRNH7g6mm62aSmA/bS+fAaFKUTXCkGWDodEIxAhwxbQOSWarIDwYpgcnJ+o/qzuh5eKwe7rv3n9C1d0NXanK6i1cWxamfJLblpzywELTUSpIktcLAzSBwYzt2vPVDzOu07XVfv785/gTmSMNNUgkAAA=="

	data, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		panic(err)
	}

	stream, err := nbt.FromBytes(data, nbt.BigEndian)
	if err != nil {
		panic(err)
	}

	tag, err := stream.ReadTag()
	if err != nil {
		panic(err)
	}

	bad, err := tag.ToString()
	if err != nil {
		panic(err)
	}

	var itemData struct{}

	if err := json.Unmarshal([]byte(bad), &itemData); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", itemData)
}
