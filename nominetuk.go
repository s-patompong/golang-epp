package main

import (
	"errors"
	"fmt"
	"github.com/qiangxue/fasthttp-routing"
	"golang-epp/src/nominetuk"
)

func nominetUkRoutes(nominet *routing.RouteGroup) {
	nominet.Get("/uk/info", nominetUkDomainInfo)
}

func nominetUkDomainInfo(c *routing.Context) error {
	domainName := string(c.FormValue("name")) // Use scottishslategifts.co.uk for testing
	if domainName == "" {
		return errors.New("No domain name")
	}

	nominetUk, _ := nominetuk.NewNominetUk()
	domain := nominetUk.NewDomain()

	info, err := domain.Info(domainName)
	if err != nil {
		return err
	}

	fmt.Fprintf(c, info.Name)

	return nil
}
