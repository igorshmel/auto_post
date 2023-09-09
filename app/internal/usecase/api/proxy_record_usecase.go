package api

import (
	"bytes"
	"context"
	"fmt"
	"github.com/igorshmel/lic_auto_post/app/internal/adapters/port"
	"github.com/igorshmel/lic_auto_post/app/pkg/config"
	"github.com/igorshmel/lic_auto_post/app/pkg/dto"
	logger "github.com/igorshmel/lic_auto_post/app/pkg/log"
	"github.com/nuttech/bell/v2"
	"io/ioutil"
	"net/http"
	"os"
)

// ProxyRecordUseCase --
type ProxyRecordUseCase struct {
	cfg             config.Config
	log             logger.Logger
	bell            *bell.Events
	persister       port.Persister
	extractor       port.Extractor
	managerDomain   port.ManagerDomain
	vkMachineDomain port.VkMachineDomain
}

// NewProxyRecordUseCase --
func NewProxyRecordUseCase(
	cfg config.Config,
	log logger.Logger,
	events *bell.Events,
	persister port.Persister,
	extractor port.Extractor,
	managerDomain port.ManagerDomain,
	vkMachineDomain port.VkMachineDomain,
) port.ProxyRecordUseCase {
	return ProxyRecordUseCase{
		cfg:             cfg,
		log:             log,
		bell:            events,
		persister:       persister,
		extractor:       extractor,
		managerDomain:   managerDomain,
		vkMachineDomain: vkMachineDomain,
	}
}

// Execute _
func (ths ProxyRecordUseCase) Execute(ctx context.Context, req *dto.ProxyRecordReqDTO) error {
	//msg := fmt.Sprintf
	log := ths.log.WithMethod("usecase ProxyRecord")

	// -- Бизнес логика --
	// ---------------------------------------------------------------------------------------------------------------------------

	// -- Инфраструктурная логика --
	// ---------------------------------------------------------------------------------------------------------------------------

	// -- Периферия --
	// ---------------------------------------------------------------------------------------------------------------------------
	proxyReq := fmt.Sprintf(`{"url": "%s","auth_url": "%s", "service": "%s"}`, req.URL, req.AuthURL, req.Service)
	// отправка данных на сервер
	jsonBody := []byte(proxyReq)
	bodyReader := bytes.NewReader(jsonBody)

	requestURL := fmt.Sprintf("http://95.163.243.113:%s%s", ths.cfg.App.Port, "/api/v1/init/")
	reqst, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}

	res, err := http.DefaultClient.Do(reqst)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	log.Info("client: got response!")
	log.Info("client: status code: %d\n", res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	log.Debug("client: response body: %s\n", resBody)

	return nil
}
