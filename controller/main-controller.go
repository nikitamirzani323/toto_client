package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"bitbucket.org/isbtotogroup/frontend_svelte/config"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

const PATH string = config.PATH_API

type clientInit struct {
	Token string `json:"token"`
}
type clientlistpasaran struct {
	Token    string `json:"token"`
	Company  string `json:"company"`
	Timezone string `json:"timezone"`
}
type clientcheckpasaran struct {
	Company      string `json:"company"`
	Timezone     string `json:"timezone"`
	Pasaran_code string `json:"pasaran_code"`
}
type clientinitpasaran struct {
	Company      string `json:"company"`
	Pasaran_code string `json:"pasaran_code"`
	Permainan    string `json:"permainan"`
}
type clientlimitpasaran struct {
	Company         string `json:"company"`
	Pasaran_code    string `json:"pasaran_code"`
	Pasaran_periode string `json:"pasaran_periode"`
	Permainan       string `json:"permainan"`
	Username        string `json:"username"`
}
type clientinvoicedetailid struct {
	Invoice   int    `json:"invoice"`
	Company   string `json:"company"`
	Username  string `json:"username"`
	Permainan string `json:"permainan"`
}
type clientresult struct {
	Company      string `json:"company"`
	Pasaran_code string `json:"pasaran_code"`
}
type clientresultall struct {
	Company string `json:"company"`
}
type clientslip struct {
	Company      string `json:"company"`
	Pasaran_code string `json:"pasaran_code"`
	Username     string `json:"username"`
}
type clientslipall struct {
	Company  string `json:"company"`
	Username string `json:"username"`
}
type clientslipdetail struct {
	Company   string `json:"company"`
	Username  string `json:"username"`
	Idinvoice string `json:"idinvoice"`
}
type clientsavetransaksi struct {
	Pasaran_idtransaction string          `json:"pasaran_idtransaction"`
	Pasaran_idcomp        string          `json:"pasaran_idcomp"`
	Token                 string          `json:"token"`
	Company               string          `json:"company"`
	Username              string          `json:"username"`
	Ipaddress             string          `json:"ipaddress"`
	Devicemember          string          `json:"devicemember"`
	Timezone              string          `json:"timezone"`
	Total                 int             `json:"total"`
	Data                  json.RawMessage `json:"data"`
}
type clientbukumimpi struct {
	Tipe string `json:"tipe"`
	Nama string `json:"nama"`
}

type responseinitresult struct {
	Status          int    `json:"status"`
	Token           string `json:"token"`
	Developer       string `json:"member_username"`
	Company         string `json:"member_company"`
	Credit          int    `json:"member_credit"`
	Website_status  string `json:"website_status"`
	Website_message string `json:"website_message"`
}
type responsecheckpasaran struct {
	Status      int    `json:"status"`
	Message     string `json:"message"`
	Totalrecord int    `json:"totalrecord"`
	Record      []struct {
		PasaranIdtransaction string `json:"pasaran_idtransaction"`
		PasaranName          string `json:"pasaran_name"`
		PasaranPeriode       string `json:"pasaran_periode"`
		PasaranIdcomp        string `json:"pasaran_idcomp"`
		PasaranStatus        string `json:"pasaran_status"`
	} `json:"record"`
	Time string `json:"time"`
}
type responseinvoicebet struct {
	Status     int         `json:"status"`
	Totalbayar int         `json:"totalbayar"`
	Totalbet   int         `json:"totalrecord"`
	Permainan  interface{} `json:"permainan"`
	Record     interface{} `json:"record"`
}
type response struct {
	Status int         `json:"status"`
	Record interface{} `json:"record"`
}
type responsedua struct {
	Status int `json:"status"`
	Record []struct {
		No      int    `json:"no"`
		Date    string `json:"date"`
		Periode string `json:"periode"`
		Result  string `json:"result"`
	} `json:"record"`
}
type responseduaall struct {
	Status int `json:"status"`
	Record []struct {
		No      int    `json:"no"`
		Date    string `json:"date"`
		Pasaran string `json:"pasaran"`
		Periode string `json:"periode"`
		Result  string `json:"result"`
	} `json:"record"`
}

func InitToken(c *fiber.Ctx) error {
	client := new(clientInit)
	render_page := time.Now()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	axios := resty.New()
	resp, err := axios.R().
		SetResult(responseinitresult{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"token": client.Token,
		}).
		Post(PATH + "api/servicetoken")
	if err != nil {
		log.Println(err.Error())
	}
	// Explore response object
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*responseinitresult)
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"status":          http.StatusOK,
		"token":           result.Token,
		"developer":       result.Developer,
		"company":         result.Company,
		"credit":          result.Credit,
		"website_status":  result.Website_status,
		"website_message": result.Website_message,
		"time":            time.Since(render_page).String(),
	})
}
func Listpasaran(c *fiber.Ctx) error {
	client := new(clientlistpasaran)
	render_page := time.Now()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	axios := resty.New()
	resp, err := axios.R().
		SetResult(response{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_company": client.Company,
		}).
		Post(PATH + "api/serviceinit")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response)
	return c.JSON(fiber.Map{
		"status": http.StatusOK,
		"record": result.Record,
		"time":   time.Since(render_page).String(),
	})
}
func Checkpasaran(c *fiber.Ctx) error {
	client := new(clientcheckpasaran)
	render_page := time.Now()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsecheckpasaran{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_company": client.Company,
			"pasaran_code":   client.Pasaran_code,
		}).
		Post(PATH + "api/servicecheckpasaran")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responsecheckpasaran)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":                http.StatusOK,
			"pasaran_idtransaction": result.Record[0].PasaranIdtransaction,
			"pasaran_name":          result.Record[0].PasaranName,
			"pasaran_periode":       result.Record[0].PasaranPeriode,
			"pasaran_idcomp":        result.Record[0].PasaranIdcomp,
			"pasaran_status":        result.Record[0].PasaranStatus,
			"time":                  time.Since(render_page).String(),
		})

	} else {
		return c.JSON(fiber.Map{
			"status": http.StatusBadRequest,
			"record": nil,
			"time":   time.Since(render_page).String(),
		})
	}
}
func Inittogel_432d(c *fiber.Ctx) error {
	client := new(clientinitpasaran)
	render_page := time.Now()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	axios := resty.New()
	resp, err := axios.R().
		SetResult(response{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_company": client.Company,
			"pasaran_code":   client.Pasaran_code,
			"permainan":      client.Permainan,
		}).
		Post(PATH + "api/serviceconfigtogel")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status": http.StatusOK,
			"record": result.Record,
			"time":   time.Since(render_page).String(),
		})
	} else {
		return c.JSON(fiber.Map{
			"status": http.StatusBadRequest,
			"record": nil,
			"time":   time.Since(render_page).String(),
		})
	}
}
func Clientlimitpasaran(c *fiber.Ctx) error {
	client := new(clientlimitpasaran)
	render_page := time.Now()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	axios := resty.New()
	resp, err := axios.R().
		SetResult(response{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_company":  client.Company,
			"client_username": client.Username,
			"pasaran_code":    client.Pasaran_code,
			"pasaran_periode": client.Pasaran_periode,
			"permainan":       client.Permainan,
		}).
		Post(PATH + "api/servicelimittogel")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status": http.StatusOK,
			"record": result.Record,
			"time":   time.Since(render_page).String(),
		})
	} else {
		return c.JSON(fiber.Map{
			"status": http.StatusOK,
			"record": nil,
			"time":   time.Since(render_page).String(),
		})
	}
}
func Resulttogel(c *fiber.Ctx) error {
	client := new(clientresult)
	render_page := time.Now()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedua{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_company": client.Company,
			"pasaran_code":   client.Pasaran_code,
		}).
		Post(PATH + "api/serviceresult")
	if err != nil {
		log.Println(err.Error())
	}

	result := resp.Result().(*responsedua)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status": http.StatusOK,
			"record": result.Record,
			"time":   time.Since(render_page).String(),
		})
	} else {
		return c.JSON(fiber.Map{
			"status": http.StatusOK,
			"record": nil,
			"time":   time.Since(render_page).String(),
		})
	}
}
func ResulttogelAll(c *fiber.Ctx) error {
	client := new(clientresultall)
	render_page := time.Now()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	axios := resty.New()
	resp, err := axios.R().
		SetResult(responseduaall{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_company": client.Company,
		}).
		Post(PATH + "api/serviceresultall")
	if err != nil {
		log.Println(err.Error())
	}

	result := resp.Result().(*responseduaall)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status": http.StatusOK,
			"record": result.Record,
			"time":   time.Since(render_page).String(),
		})
	} else {
		return c.JSON(fiber.Map{
			"status": http.StatusOK,
			"record": nil,
			"time":   time.Since(render_page).String(),
		})
	}
}
func Invoicebet(c *fiber.Ctx) error {
	client := new(clientlimitpasaran)
	render_page := time.Now()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	axios := resty.New()
	resp, err := axios.R().
		SetResult(responseinvoicebet{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_company":  client.Company,
			"client_username": client.Username,
			"pasaran_code":    client.Pasaran_code,
			"pasaran_periode": client.Pasaran_periode,
		}).
		Post(PATH + "api/serviceinvoicebet")
	if err != nil {
		log.Println(err.Error())
	}
	// log.Println(resp)
	result := resp.Result().(*responseinvoicebet)
	// log.Println(resp)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":     http.StatusOK,
			"totalbayar": result.Totalbayar,
			"totalbet":   result.Totalbet,
			"record":     result.Record,
			"time":       time.Since(render_page).String(),
		})
	} else {
		return c.JSON(fiber.Map{
			"status": http.StatusOK,
			"record": nil,
			"time":   time.Since(render_page).String(),
		})
	}
}
func Invoicebetid(c *fiber.Ctx) error {
	client := new(clientinvoicedetailid)
	render_page := time.Now()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	axios := resty.New()
	resp, err := axios.R().
		SetResult(response{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_idinvoice": client.Invoice,
			"client_company":   client.Company,
			"client_username":  client.Username,
			"permainan":        client.Permainan,
		}).
		Post(PATH + "api/serviceinvoicebetdetail")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status": http.StatusOK,
			"record": result.Record,
			"time":   time.Since(render_page).String(),
		})
	} else {
		return c.JSON(fiber.Map{
			"status": http.StatusOK,
			"record": nil,
			"time":   time.Since(render_page).String(),
		})
	}
}
func Slipperiode(c *fiber.Ctx) error {
	client := new(clientslip)
	render_page := time.Now()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	axios := resty.New()
	resp, err := axios.R().
		SetResult(response{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_company":  client.Company,
			"client_username": client.Username,
			"pasaran_code":    client.Pasaran_code,
		}).
		Post(PATH + "api/serviceslip")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status": http.StatusOK,
			"record": result.Record,
			"time":   time.Since(render_page).String(),
		})
	} else {
		return c.JSON(fiber.Map{
			"status": http.StatusOK,
			"record": nil,
			"time":   time.Since(render_page).String(),
		})
	}
}
func SlipperiodeAll(c *fiber.Ctx) error {
	client := new(clientslipall)
	render_page := time.Now()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	axios := resty.New()
	resp, err := axios.R().
		SetResult(response{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_company":  client.Company,
			"client_username": client.Username,
		}).
		Post(PATH + "api/serviceslipall")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status": http.StatusOK,
			"record": result.Record,
			"time":   time.Since(render_page).String(),
		})
	} else {
		return c.JSON(fiber.Map{
			"status": http.StatusOK,
			"record": nil,
			"time":   time.Since(render_page).String(),
		})
	}
}
func Slipperiodedetail(c *fiber.Ctx) error {
	client := new(clientslipdetail)
	render_page := time.Now()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	axios := resty.New()
	resp, err := axios.R().
		SetResult(response{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_company":  client.Company,
			"client_username": client.Username,
			"idtrxkeluaran":   client.Idinvoice,
		}).
		Post(PATH + "api/serviceslipdetail")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(resp)
	result := resp.Result().(*response)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status": http.StatusOK,
			"record": result.Record,
			"time":   time.Since(render_page).String(),
		})
	} else {
		return c.JSON(fiber.Map{
			"status": http.StatusOK,
			"record": nil,
			"time":   time.Since(render_page).String(),
		})
	}
}
func Savetransaksi(c *fiber.Ctx) error {
	client := new(clientsavetransaksi)
	render_page := time.Now()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	axios := resty.New()

	resp, err := axios.R().
		SetResult(response{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"idtrxkeluaran":   client.Pasaran_idtransaction,
			"idcomppasaran":   client.Pasaran_idcomp,
			"devicemember":    client.Devicemember,
			"formipaddress":   client.Ipaddress,
			"timezone":        client.Timezone,
			"client_company":  client.Company,
			"client_username": client.Username,
			"totalbayarbet":   client.Total,
			"list4d":          string(client.Data),
		}).
		Post(PATH + "api/savetransaksi")
	if err != nil {
		log.Println(err.Error())
	}
	// Explore response object
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*response)

	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status": http.StatusOK,
			"record": result.Record,
			"time":   time.Since(render_page).String(),
		})
	} else {
		return c.JSON(fiber.Map{
			"status": http.StatusOK,
			"record": nil,
			"time":   time.Since(render_page).String(),
		})
	}
}
func Listbukumimpi(c *fiber.Ctx) error {
	client := new(clientbukumimpi)
	render_page := time.Now()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	axios := resty.New()
	resp, err := axios.R().
		SetResult(response{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"tipe": client.Tipe,
			"nama": client.Nama,
		}).
		Post(PATH + "api/bukumimpi")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*response)

	return c.JSON(fiber.Map{
		"status": http.StatusOK,
		"record": result.Record,
		"time":   time.Since(render_page).String(),
	})
}
