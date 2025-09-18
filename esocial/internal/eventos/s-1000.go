// S-1000 - Informações do Empregador/Contribuinte/ÓrgãoPúblico
package eventos

import (
	"encoding/xml"

	"github.com/avelinobego/esocial/internal/tipos"
)

type S1000 struct {
	XMLName  xml.Name        `xml:"evtInfoEmpregador"`
	Id       tipos.Id        `xml:"Id,attr"`
	Idevento tipos.IdeEvento `xml:"ideEvento"`
}
