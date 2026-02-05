package entity

import "encoding/xml"

type S1000 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Esocial struct {
	Evtinfoempregador   Evtinfoempregador   `xml:"evtInfoEmpregador,omitempty"`
	Evttabestab         Evttabestab         `xml:"evtTabEstab,omitempty"`
	Evttabrubrica       Evttabrubrica       `xml:"evtTabRubrica,omitempty"`
	Evttablotacao       Evttablotacao       `xml:"evtTabLotacao,omitempty"`
	Evttabprocesso      Evttabprocesso      `xml:"evtTabProcesso,omitempty"`
	Evtremun            Evtremun            `xml:"evtRemun,omitempty"`
	Evtrmnrpps          Evtrmnrpps          `xml:"evtRmnRPPS,omitempty"`
	Evtbenprrp          Evtbenprrp          `xml:"evtBenPrRP,omitempty"`
	Evtpgtos            Evtpgtos            `xml:"evtPgtos,omitempty"`
	Evtcomprod          Evtcomprod          `xml:"evtComProd,omitempty"`
	Evtcontratavnp      Evtcontratavnp      `xml:"evtContratAvNP,omitempty"`
	Evtinfocomplper     Evtinfocomplper     `xml:"evtInfoComplPer,omitempty"`
	Evtreabreevper      Evtreabreevper      `xml:"evtReabreEvPer,omitempty"`
	Evtfechaevper       Evtfechaevper       `xml:"evtFechaEvPer,omitempty"`
	Evtadmprelim        Evtadmprelim        `xml:"evtAdmPrelim,omitempty"`
	Evtadmissao         Evtadmissao         `xml:"evtAdmissao,omitempty"`
	Evtaltcadastral     Evtaltcadastral     `xml:"evtAltCadastral,omitempty"`
	Evtaltcontratual    Evtaltcontratual    `xml:"evtAltContratual,omitempty"`
	Evtcat              Evtcat              `xml:"evtCAT,omitempty"`
	Evtmonit            Evtmonit            `xml:"evtMonit,omitempty"`
	Evttoxic            Evttoxic            `xml:"evtToxic,omitempty"`
	Evtafasttemp        Evtafasttemp        `xml:"evtAfastTemp,omitempty"`
	Evtcessao           Evtcessao           `xml:"evtCessao,omitempty"`
	Evtexprisco         Evtexprisco         `xml:"evtExpRisco,omitempty"`
	Evtreintegr         Evtreintegr         `xml:"evtReintegr,omitempty"`
	Evtdeslig           Evtdeslig           `xml:"evtDeslig,omitempty"`
	Evttsvinicio        Evttsvinicio        `xml:"evtTSVInicio,omitempty"`
	Evttsvaltcontr      Evttsvaltcontr      `xml:"evtTSVAltContr,omitempty"`
	Evttsvtermino       Evttsvtermino       `xml:"evtTSVTermino,omitempty"`
	Evtcdbenefin        Evtcdbenefin        `xml:"evtCdBenefIn,omitempty"`
	Evtcdbenefalt       Evtcdbenefalt       `xml:"evtCdBenefAlt,omitempty"`
	Evtcdbenin          Evtcdbenin          `xml:"evtCdBenIn,omitempty"`
	Evtcdbenalt         Evtcdbenalt         `xml:"evtCdBenAlt,omitempty"`
	Evtreativben        Evtreativben        `xml:"evtReativBen,omitempty"`
	Evtcdbenterm        Evtcdbenterm        `xml:"evtCdBenTerm,omitempty"`
	Evtproctrab         Evtproctrab         `xml:"evtProcTrab,omitempty"`
	Evtcontproc         Evtcontproc         `xml:"evtContProc,omitempty"`
	Evtconsolidcontproc Evtconsolidcontproc `xml:"evtConsolidContProc,omitempty"`
	Evtexclusao         Evtexclusao         `xml:"evtExclusao,omitempty"`
	Evtexcproctrab      Evtexcproctrab      `xml:"evtExcProcTrab,omitempty"`
	Evtbasestrab        Evtbasestrab        `xml:"evtBasesTrab,omitempty"`
	Evtirrfbenef        Evtirrfbenef        `xml:"evtIrrfBenef,omitempty"`
	Evtbasesfgts        Evtbasesfgts        `xml:"evtBasesFGTS,omitempty"`
	Evtcs               Evtcs               `xml:"evtCS,omitempty"`
	Evtirrf             Evtirrf             `xml:"evtIrrf,omitempty"`
	Evtfgts             Evtfgts             `xml:"evtFGTS,omitempty"`
	Evttribproctrab     Evttribproctrab     `xml:"evtTribProcTrab,omitempty"`
	Evtfgtsproctrab     Evtfgtsproctrab     `xml:"evtFGTSProcTrab,omitempty"`
	Evtanotjud          Evtanotjud          `xml:"evtAnotJud,omitempty"`
	Evtbaixa            Evtbaixa            `xml:"evtBaixa,omitempty"`
}

type Evtinfoempregador struct {
	Id             string        `xml:"Id"`
	Ideevento      Ideevento     `xml:"ideEvento"`
	Ideempregador  Ideempregador `xml:"ideEmpregador"`
	Infoempregador string        `xml:"infoEmpregador"`
}

type Ideevento struct {
	Tpamb   int64  `xml:"tpAmb"`
	Procemi int64  `xml:"procEmi"`
	Verproc string `xml:"verProc"`
}

type Ideempregador struct {
	Tpinsc int64  `xml:"tpInsc"`
	Nrinsc string `xml:"nrInsc"`
}

type Infoempregador struct {
	Inclusao  []Inclusao  `xml:"inclusao,omitempty"`
	Alteracao []Alteracao `xml:"alteracao,omitempty"`
	Exclusao  []Exclusao  `xml:"exclusao,omitempty"`
}

type Inclusao struct {
	Ideperiodo   Ideperiodo     `xml:"idePeriodo"`
	Infocadastro []Infocadastro `xml:"infoCadastro,omitempty"`
	Ideestab     Ideestab       `xml:"ideEstab"`
	Dadosestab   Dadosestab     `xml:"dadosEstab"`
	Iderubrica   Iderubrica     `xml:"ideRubrica"`
	Dadosrubrica Dadosrubrica   `xml:"dadosRubrica"`
	Idelotacao   Idelotacao     `xml:"ideLotacao"`
	Dadoslotacao Dadoslotacao   `xml:"dadosLotacao"`
	Ideprocesso  Ideprocesso    `xml:"ideProcesso"`
	Dadosproc    Dadosproc      `xml:"dadosProc"`
}

type Ideperiodo struct {
	Inivalid string `xml:"iniValid"`
	Fimvalid string `xml:"fimValid,omitempty"`
}

type Infocadastro struct {
	Classtrib            string                 `xml:"classTrib"`
	Indcoop              []string               `xml:"indCoop,omitempty"`
	Indconstr            []string               `xml:"indConstr,omitempty"`
	Inddesfolha          int64                  `xml:"indDesFolha"`
	Indopccp             []string               `xml:"indOpcCP,omitempty"`
	Indporte             []string               `xml:"indPorte,omitempty"`
	Indoptregeletron     int64                  `xml:"indOptRegEletron"`
	Cnpjefr              []string               `xml:"cnpjEFR,omitempty"`
	Dttrans11096         []string               `xml:"dtTrans11096,omitempty"`
	Indtribfolhapispasep []string               `xml:"indTribFolhaPisPasep,omitempty"`
	Indpertirrf          []string               `xml:"indPertIRRF,omitempty"`
	Dadosisencao         []Dadosisencao         `xml:"dadosIsencao,omitempty"`
	Infoorginternacional []Infoorginternacional `xml:"infoOrgInternacional,omitempty"`
}

type Dadosisencao struct {
	Ideminlei    string   `xml:"ideMinLei"`
	Nrcertif     string   `xml:"nrCertif"`
	Dtemiscertif string   `xml:"dtEmisCertif"`
	Dtvenccertif string   `xml:"dtVencCertif"`
	Nrprotrenov  []string `xml:"nrProtRenov,omitempty"`
	Dtprotrenov  []string `xml:"dtProtRenov,omitempty"`
	Dtdou        []string `xml:"dtDou,omitempty"`
	Pagdou       []string `xml:"pagDou,omitempty"`
}

type Infoorginternacional struct {
	Indacordoisenmulta int64 `xml:"indAcordoIsenMulta"`
}

type Alteracao struct {
	Ideperiodo Ideperiodo `xml:"idePeriodo"`
}

type Exclusao struct {
	Ideperiodo  Ideperiodo  `xml:"idePeriodo"`
	Ideestab    Ideestab    `xml:"ideEstab"`
	Iderubrica  Iderubrica  `xml:"ideRubrica"`
	Idelotacao  Idelotacao  `xml:"ideLotacao"`
	Ideprocesso Ideprocesso `xml:"ideProcesso"`
}

type S1005 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evttabestab struct {
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Infoestab     string        `xml:"infoEstab"`
}

type Infoestab struct {
	Inclusao      []Inclusao      `xml:"inclusao,omitempty"`
	Alteracao     []Alteracao     `xml:"alteracao,omitempty"`
	Exclusao      []Exclusao      `xml:"exclusao,omitempty"`
	Cnaeprep      int64           `xml:"cnaePrep"`
	Cnpjresp      []string        `xml:"cnpjResp,omitempty"`
	Aliqrat       int64           `xml:"aliqRat"`
	Fap           []string        `xml:"fap,omitempty"`
	Aliqratajust  []string        `xml:"aliqRatAjust,omitempty"`
	Infoestabref  []Infoestabref  `xml:"infoEstabRef,omitempty"`
	Infocomplobra []Infocomplobra `xml:"infoComplObra,omitempty"`
}

type Ideestab struct {
	Tpinsc   int64    `xml:"tpInsc"`
	Nrinsc   string   `xml:"nrInsc"`
	Inivalid string   `xml:"iniValid"`
	Fimvalid []string `xml:"fimValid,omitempty"`
}

type Dadosestab struct {
	Cnaeprep   int64        `xml:"cnaePrep"`
	Cnpjresp   []string     `xml:"cnpjResp,omitempty"`
	Aliqgilrat []Aliqgilrat `xml:"aliqGilrat,omitempty"`
	Infocaepf  []Infocaepf  `xml:"infoCaepf,omitempty"`
	Infoobra   []Infoobra   `xml:"infoObra,omitempty"`
	Infotrab   []Infotrab   `xml:"infoTrab,omitempty"`
}

type Aliqgilrat struct {
	Aliqrat       []string        `xml:"aliqRat,omitempty"`
	Fap           []string        `xml:"fap,omitempty"`
	Procadmjudrat []Procadmjudrat `xml:"procAdmJudRat,omitempty"`
	Procadmjudfap []Procadmjudfap `xml:"procAdmJudFap,omitempty"`
}

type Procadmjudrat struct {
	Tpproc  int64  `xml:"tpProc"`
	Nrproc  string `xml:"nrProc"`
	Codsusp int64  `xml:"codSusp"`
}

type Procadmjudfap struct {
	Tpproc  int64  `xml:"tpProc"`
	Nrproc  string `xml:"nrProc"`
	Codsusp int64  `xml:"codSusp"`
}

type Infocaepf struct {
	Tpcaepf int64 `xml:"tpCaepf"`
}

type Infoobra struct {
	Indsubstpatrobra int64 `xml:"indSubstPatrObra"`
}

type Infotrab struct {
	Infoapr []Infoapr `xml:"infoApr,omitempty"`
	Infopcd []Infopcd `xml:"infoPCD,omitempty"`
}

type Infoapr struct {
	Nrprocjud   []string      `xml:"nrProcJud,omitempty"`
	Infoenteduc []Infoenteduc `xml:"infoEntEduc"`
}

type Infoenteduc struct {
	Nrinsc string `xml:"nrInsc"`
}

type Infopcd struct {
	Nrprocjud string `xml:"nrProcJud"`
}

type Novavalidade struct {
	Inivalid string   `xml:"iniValid"`
	Fimvalid []string `xml:"fimValid,omitempty"`
}

type S1010 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evttabrubrica struct {
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Inforubrica   string        `xml:"infoRubrica"`
}

type Inforubrica struct {
	Inclusao  []Inclusao  `xml:"inclusao,omitempty"`
	Alteracao []Alteracao `xml:"alteracao,omitempty"`
	Exclusao  []Exclusao  `xml:"exclusao,omitempty"`
}

type Iderubrica struct {
	Codrubr    string   `xml:"codRubr"`
	Idetabrubr string   `xml:"ideTabRubr"`
	Inivalid   string   `xml:"iniValid"`
	Fimvalid   []string `xml:"fimValid,omitempty"`
}

type Dadosrubrica struct {
	Dscrubr             string                `xml:"dscRubr"`
	Natrubr             int64                 `xml:"natRubr"`
	Tprubr              int64                 `xml:"tpRubr"`
	Codinccp            string                `xml:"codIncCP"`
	Codincirrf          int64                 `xml:"codIncIRRF"`
	Codincfgts          string                `xml:"codIncFGTS"`
	Codinccprp          []string              `xml:"codIncCPRP,omitempty"`
	Codincpispasep      []string              `xml:"codIncPisPasep,omitempty"`
	Tetoremun           []string              `xml:"tetoRemun,omitempty"`
	Observacao          []string              `xml:"observacao,omitempty"`
	Ideprocessocp       []Ideprocessocp       `xml:"ideProcessoCP"`
	Ideprocessoirrf     []Ideprocessoirrf     `xml:"ideProcessoIRRF"`
	Ideprocessofgts     []Ideprocessofgts     `xml:"ideProcessoFGTS"`
	Ideprocessopispasep []Ideprocessopispasep `xml:"ideProcessoPisPasep"`
}

type Ideprocessocp struct {
	Tpproc     int64  `xml:"tpProc"`
	Nrproc     string `xml:"nrProc"`
	Extdecisao int64  `xml:"extDecisao"`
	Codsusp    int64  `xml:"codSusp"`
}

type Ideprocessoirrf struct {
	Nrproc  string `xml:"nrProc"`
	Codsusp int64  `xml:"codSusp"`
}

type Ideprocessofgts struct {
	Nrproc string `xml:"nrProc"`
}

type Ideprocessopispasep struct {
	Nrproc  string `xml:"nrProc"`
	Codsusp int64  `xml:"codSusp"`
}

type S1020 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evttablotacao struct {
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Infolotacao   string        `xml:"infoLotacao"`
}

type Infolotacao struct {
	Inclusao  []Inclusao  `xml:"inclusao,omitempty"`
	Alteracao []Alteracao `xml:"alteracao,omitempty"`
	Exclusao  []Exclusao  `xml:"exclusao,omitempty"`
}

type Idelotacao struct {
	Inivalid string `xml:"iniValid"`
	Fimvalid string `xml:"fimValid,omitempty"`
}

type Dadoslotacao struct {
	Tplotacao       string            `xml:"tpLotacao"`
	Tpinsc          []string          `xml:"tpInsc,omitempty"`
	Nrinsc          []string          `xml:"nrInsc,omitempty"`
	Fpaslotacao     Fpaslotacao       `xml:"fpasLotacao"`
	Infoemprparcial []Infoemprparcial `xml:"infoEmprParcial,omitempty"`
	Dadosopport     []Dadosopport     `xml:"dadosOpPort,omitempty"`
}

type Fpaslotacao struct {
	Fpas                 int64                  `xml:"fpas"`
	Codtercs             string                 `xml:"codTercs"`
	Codtercssusp         []string               `xml:"codTercsSusp,omitempty"`
	Infoprocjudterceiros []Infoprocjudterceiros `xml:"infoProcJudTerceiros,omitempty"`
}

type Infoprocjudterceiros struct {
	Procjudterceiro []Procjudterceiro `xml:"procJudTerceiro"`
}

type Procjudterceiro struct {
	Codterc   string `xml:"codTerc"`
	Nrprocjud string `xml:"nrProcJud"`
	Codsusp   int64  `xml:"codSusp"`
}

type Infoemprparcial struct {
	Tpinsccontrat int64    `xml:"tpInscContrat"`
	Nrinsccontrat string   `xml:"nrInscContrat"`
	Tpinscprop    []string `xml:"tpInscProp,omitempty"`
	Nrinscprop    []string `xml:"nrInscProp,omitempty"`
}

type Dadosopport struct {
	Aliqrat         int64   `xml:"aliqRat"`
	Fap             float64 `xml:"fap"`
	Cnpjopportuario string  `xml:"cnpjOpPortuario"`
}

type S1070 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evttabprocesso struct {
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Infoprocesso  string        `xml:"infoProcesso"`
}

type Infoprocesso struct {
	Inclusao    []Inclusao  `xml:"inclusao,omitempty"`
	Alteracao   []Alteracao `xml:"alteracao,omitempty"`
	Exclusao    []Exclusao  `xml:"exclusao,omitempty"`
	Origem      int64       `xml:"origem"`
	Obsproctrab []string    `xml:"obsProcTrab,omitempty"`
	Dadoscompl  string      `xml:"dadosCompl"`
	Nrproctrab  string      `xml:"nrProcTrab"`
	Dtsent      string      `xml:"dtSent"`
	Ufvara      string      `xml:"ufVara"`
	Codmunic    int64       `xml:"codMunic"`
	Idvara      int64       `xml:"idVara"`
}

type Ideprocesso struct {
	Tpproc   int64    `xml:"tpProc"`
	Nrproc   string   `xml:"nrProc"`
	Inivalid string   `xml:"iniValid"`
	Fimvalid []string `xml:"fimValid,omitempty"`
}

type Dadosproc struct {
	Indautoria   []string       `xml:"indAutoria,omitempty"`
	Indmatproc   int64          `xml:"indMatProc"`
	Observacao   []string       `xml:"observacao,omitempty"`
	Dadosprocjud []Dadosprocjud `xml:"dadosProcJud,omitempty"`
	Infosusp     []Infosusp     `xml:"infoSusp"`
}

type Dadosprocjud struct {
	Ufvara   string `xml:"ufVara"`
	Codmunic int64  `xml:"codMunic"`
	Idvara   int64  `xml:"idVara"`
}

type Infosusp struct {
	Codsusp     int64  `xml:"codSusp"`
	Indsusp     string `xml:"indSusp"`
	Dtdecisao   string `xml:"dtDecisao"`
	Inddeposito string `xml:"indDeposito"`
}

type S1200 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
	string  `xml:""`
}

type Evtremun struct {
	Id             string         `xml:"Id"`
	Ideevento      Ideevento      `xml:"ideEvento"`
	Ideempregador  Ideempregador  `xml:"ideEmpregador"`
	Idetrabalhador Idetrabalhador `xml:"ideTrabalhador"`
	Dmdev          []Dmdev        `xml:"dmDev"`
}

type Idetrabalhador struct {
	Infomv        []Infomv        `xml:"infoMV,omitempty"`
	Infointerm    []Infointerm    `xml:"infoInterm"`
	Infocomplem   []Infocomplem   `xml:"infoComplem,omitempty"`
	Infocompl     []Infocompl     `xml:"infoCompl,omitempty"`
	Procjudtrab   []Procjudtrab   `xml:"procJudTrab"`
	Cpfbenef      string          `xml:"cpfBenef"`
	Dmdev         []Dmdev         `xml:"dmDev"`
	Totinfoir     []Totinfoir     `xml:"totInfoIR,omitempty"`
	Infoircomplem []Infoircomplem `xml:"infoIRComplem"`
	Cpftrab       string          `xml:"cpfTrab"`
}

type Infomv struct {
	Indmv         int64           `xml:"indMV"`
	Remunoutrempr []Remunoutrempr `xml:"remunOutrEmpr"`
}

type Remunoutrempr struct {
	Tpinsc     int64   `xml:"tpInsc"`
	Nrinsc     string  `xml:"nrInsc"`
	Codcateg   int64   `xml:"codCateg"`
	Vlrremunoe float64 `xml:"vlrRemunOE"`
}

type Infocomplem struct {
	Nmtrab       string         `xml:"nmTrab"`
	Dtnascto     string         `xml:"dtNascto"`
	Sucessaovinc []Sucessaovinc `xml:"sucessaoVinc,omitempty"`
}

type Sucessaovinc struct {
	Cnpjorgaoant string   `xml:"cnpjOrgaoAnt"`
	Dtexercicio  string   `xml:"dtExercicio"`
	Observacao   []string `xml:"observacao,omitempty"`
	Dttransf     string   `xml:"dtTransf"`
	Tpinsc       int64    `xml:"tpInsc"`
	Nrinsc       string   `xml:"nrInsc"`
	Matricant    []string `xml:"matricAnt,omitempty"`
	Dtadm        string   `xml:"dtAdm"`
}

type Procjudtrab struct {
	Tptrib    int64  `xml:"tpTrib"`
	Nrprocjud string `xml:"nrProcJud"`
	Codsusp   int64  `xml:"codSusp"`
}

type Infointerm struct {
	Dia     int64    `xml:"dia"`
	Hrstrab []string `xml:"hrsTrab,omitempty"`
}

type Dmdev struct {
	Infocomplcont []Infocomplcont `xml:"infoComplCont,omitempty"`
	Nrbeneficio   string          `xml:"nrBeneficio"`
	Indrra        []string        `xml:"indRRA,omitempty"`
	Notaft        []string        `xml:"notAFT,omitempty"`
	Infoperapur   []Infoperapur   `xml:"infoPerApur,omitempty"`
	Infoperant    []Infoperant    `xml:"infoPerAnt,omitempty"`
	Idedmdev      string          `xml:"ideDmDev"`
	Tppgto        int64           `xml:"tpPgto"`
	Dtpgto        string          `xml:"dtPgto"`
	Codcateg      int64           `xml:"codCateg"`
	Infoir        []Infoir        `xml:"infoIR"`
	Totapurmen    []Totapurmen    `xml:"totApurMen"`
	Totapurdia    []Totapurdia    `xml:"totApurDia"`
	Inforra       []Inforra       `xml:"infoRRA,omitempty"`
	Infopgtoext   []Infopgtoext   `xml:"infoPgtoExt,omitempty"`
}

type Inforra struct {
	Tpprocrra   int64         `xml:"tpProcRRA"`
	Nrprocrra   []string      `xml:"nrProcRRA,omitempty"`
	Descrra     string        `xml:"descRRA"`
	Qtdmesesrra float64       `xml:"qtdMesesRRA"`
	Despprocjud []Despprocjud `xml:"despProcJud,omitempty"`
	Ideadv      []Ideadv      `xml:"ideAdv"`
}

type Despprocjud struct {
	Vlrdespcustas    float64 `xml:"vlrDespCustas"`
	Vlrdespadvogados float64 `xml:"vlrDespAdvogados"`
}

type Ideadv struct {
	Tpinsc int64    `xml:"tpInsc"`
	Nrinsc string   `xml:"nrInsc"`
	Vlradv []string `xml:"vlrAdv,omitempty"`
}

type Infoperapur struct {
	Ideestab    []Ideestab    `xml:"ideEstab"`
	Ideestablot []Ideestablot `xml:"ideEstabLot"`
}

type Ideestablot struct {
	Qtddiasav      []string         `xml:"qtdDiasAv,omitempty"`
	Remunperapur   []Remunperapur   `xml:"remunPerApur"`
	Remunperant    []Remunperant    `xml:"remunPerAnt"`
	Infoagnocivo   []Infoagnocivo   `xml:"infoAgNocivo,omitempty"`
	Detverbas      []Detverbas      `xml:"detVerbas"`
	Infosimples    []Infosimples    `xml:"infoSimples,omitempty"`
	Tpinsc         int64            `xml:"tpInsc"`
	Nrinsc         string           `xml:"nrInsc"`
	Codlotacao     string           `xml:"codLotacao"`
	Infocategincid []Infocategincid `xml:"infoCategIncid"`
}

type Remunperapur struct {
	Indsimples   []string       `xml:"indSimples,omitempty"`
	Infoagnocivo []Infoagnocivo `xml:"infoAgNocivo,omitempty"`
	Matricula    []string       `xml:"matricula,omitempty"`
	Itensremun   []Itensremun   `xml:"itensRemun"`
}

type Itensremun struct {
	Descfolha  []Descfolha `xml:"descFolha,omitempty"`
	Codrubr    string      `xml:"codRubr"`
	Idetabrubr string      `xml:"ideTabRubr"`
	Qtdrubr    []string    `xml:"qtdRubr,omitempty"`
	Fatorrubr  []string    `xml:"fatorRubr,omitempty"`
	Vrrubr     float64     `xml:"vrRubr"`
	Indapurir  int64       `xml:"indApurIR"`
}

type Descfolha struct {
	Tpdesc     int64    `xml:"tpDesc"`
	Instfinanc string   `xml:"instFinanc"`
	Nrdoc      string   `xml:"nrDoc"`
	Observacao []string `xml:"observacao,omitempty"`
}

type Infoagnocivo struct {
	Grauexp int64 `xml:"grauExp"`
}

type Infoperant struct {
	Remunorgsuc string       `xml:"remunOrgSuc"`
	Ideperiodo  []Ideperiodo `xml:"idePeriodo"`
	Ideadc      []Ideadc     `xml:"ideADC"`
}

type Ideadc struct {
	Ideperiodo []Ideperiodo `xml:"idePeriodo"`
	Dtacconv   []string     `xml:"dtAcConv,omitempty"`
	Tpacconv   string       `xml:"tpAcConv"`
	Dsc        string       `xml:"dsc"`
	Remunsuc   []string     `xml:"remunSuc,omitempty"`
}

type Remunperant struct {
	Indsimples   []string       `xml:"indSimples,omitempty"`
	Infoagnocivo []Infoagnocivo `xml:"infoAgNocivo,omitempty"`
	Matricula    []string       `xml:"matricula,omitempty"`
	Itensremun   []Itensremun   `xml:"itensRemun"`
}

type Infocomplcont struct {
	Codcbo       string   `xml:"codCBO"`
	Natatividade []string `xml:"natAtividade,omitempty"`
	Qtddiastrab  []string `xml:"qtdDiasTrab,omitempty"`
}

type S1202 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evtrmnrpps struct {
	Id             string         `xml:"Id"`
	Ideevento      Ideevento      `xml:"ideEvento"`
	Ideempregador  Ideempregador  `xml:"ideEmpregador"`
	Idetrabalhador Idetrabalhador `xml:"ideTrabalhador"`
	Dmdev          []Dmdev        `xml:"dmDev"`
}

type S1207 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evtbenprrp struct {
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Idebenef      Idebenef      `xml:"ideBenef"`
	Dmdev         []Dmdev       `xml:"dmDev"`
}

type Idebenef struct {
	Infopgto      []Infopgto      `xml:"infoPgto"`
	Infoircomplem []Infoircomplem `xml:"infoIRComplem"`
	Cpfbenef      string          `xml:"cpfBenef"`
}

type S1210 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
	string  `xml:""`
}

type Evtpgtos struct {
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Idebenef      Idebenef      `xml:"ideBenef"`
}

type Infopgto struct {
	Dtpgto       string        `xml:"dtPgto"`
	Tppgto       int64         `xml:"tpPgto"`
	Perref       string        `xml:"perRef"`
	Idedmdev     string        `xml:"ideDmDev"`
	Vrliq        float64       `xml:"vrLiq"`
	Paisresidext []string      `xml:"paisResidExt,omitempty"`
	Infopgtoext  []Infopgtoext `xml:"infoPgtoExt,omitempty"`
}

type Infopgtoext struct {
	Paisresidext string   `xml:"paisResidExt"`
	Indnif       int64    `xml:"indNIF"`
	Nifbenef     []string `xml:"nifBenef,omitempty"`
	Frmtribut    string   `xml:"frmTribut"`
	Endext       []Endext `xml:"endExt,omitempty"`
}

type Endext struct {
	Enddsclograd []string `xml:"endDscLograd,omitempty"`
	Endnrlograd  []string `xml:"endNrLograd,omitempty"`
	Endcomplem   []string `xml:"endComplem,omitempty"`
	Endbairro    []string `xml:"endBairro,omitempty"`
	Endcidade    []string `xml:"endCidade,omitempty"`
	Endestado    []string `xml:"endEstado,omitempty"`
	Endcodpostal []string `xml:"endCodPostal,omitempty"`
	Telef        []string `xml:"telef,omitempty"`
}

type Infoircomplem struct {
	Infodep      []Infodep      `xml:"infoDep"`
	Dtlaudo      []string       `xml:"dtLaudo,omitempty"`
	Perant       []Perant       `xml:"perAnt,omitempty"`
	Idedep       []Idedep       `xml:"ideDep"`
	Infoircr     []Infoircr     `xml:"infoIRCR"`
	Plansaude    []Plansaude    `xml:"planSaude"`
	Inforeembmed []Inforeembmed `xml:"infoReembMed"`
}

type Perant struct {
	Perrefajuste  string `xml:"perRefAjuste"`
	Nrrec1210orig string `xml:"nrRec1210Orig"`
}

type Infodep struct {
	Cpfdep   string   `xml:"cpfDep"`
	Dtnascto []string `xml:"dtNascto,omitempty"`
	Nome     []string `xml:"nome,omitempty"`
	Depirrf  []string `xml:"depIRRF,omitempty"`
	Tpdep    []string `xml:"tpDep,omitempty"`
	Descrdep []string `xml:"descrDep,omitempty"`
}

type Infoircr struct {
	Tpcr        string        `xml:"tpCR"`
	Deddepen    []Deddepen    `xml:"dedDepen"`
	Penalim     []Penalim     `xml:"penAlim"`
	Previdcompl []Previdcompl `xml:"previdCompl"`
	Infoprocret []Infoprocret `xml:"infoProcRet"`
}

type Deddepen struct {
	Vlrdeducao float64 `xml:"vlrDeducao"`
	Tprend     int64   `xml:"tpRend"`
	Cpfdep     string  `xml:"cpfDep"`
	Vlrdeddep  float64 `xml:"vlrDedDep"`
}

type Penalim struct {
	Vlrpensao     float64 `xml:"vlrPensao"`
	Tprend        int64   `xml:"tpRend"`
	Cpfdep        string  `xml:"cpfDep"`
	Vlrdedpenalim float64 `xml:"vlrDedPenAlim"`
}

type Previdcompl struct {
	Tpprev          int64    `xml:"tpPrev"`
	Cnpjentidpc     string   `xml:"cnpjEntidPC"`
	Vlrdedpc        []string `xml:"vlrDedPC,omitempty"`
	Vlrdedpc13      []string `xml:"vlrDedPC13,omitempty"`
	Vlrpatrocfunp   []string `xml:"vlrPatrocFunp,omitempty"`
	Vlrpatrocfunp13 []string `xml:"vlrPatrocFunp13,omitempty"`
}

type Infoprocret struct {
	Tpprocret   int64         `xml:"tpProcRet"`
	Nrprocret   string        `xml:"nrProcRet"`
	Codsusp     []string      `xml:"codSusp,omitempty"`
	Infovalores []Infovalores `xml:"infoValores"`
}

type Infovalores struct {
	Indapuracao  int64     `xml:"indApuracao"`
	Vlrnretido   []string  `xml:"vlrNRetido,omitempty"`
	Vlrdepjud    []string  `xml:"vlrDepJud,omitempty"`
	Vlrcmpanocal []string  `xml:"vlrCmpAnoCal,omitempty"`
	Vlrcmpanoant []string  `xml:"vlrCmpAnoAnt,omitempty"`
	Vlrrendsusp  []string  `xml:"vlrRendSusp,omitempty"`
	Dedsusp      []Dedsusp `xml:"dedSusp"`
}

type Dedsusp struct {
	Indtpdeducao  int64      `xml:"indTpDeducao"`
	Vlrdedsusp    []string   `xml:"vlrDedSusp,omitempty"`
	Cnpjentidpc   []string   `xml:"cnpjEntidPC,omitempty"`
	Vlrpatrocfunp []string   `xml:"vlrPatrocFunp,omitempty"`
	Benefpen      []Benefpen `xml:"benefPen"`
}

type Benefpen struct {
	Cpfdep       string  `xml:"cpfDep"`
	Vlrdepensusp float64 `xml:"vlrDepenSusp"`
}

type Plansaude struct {
	Cnpjoper    string       `xml:"cnpjOper"`
	Regans      []string     `xml:"regANS,omitempty"`
	Vlrsaudetit float64      `xml:"vlrSaudeTit"`
	Infodepsau  []Infodepsau `xml:"infoDepSau"`
}

type Infodepsau struct {
	Cpfdep      string  `xml:"cpfDep"`
	Vlrsaudedep float64 `xml:"vlrSaudeDep"`
}

type Inforeembmed struct {
	Indorgreemb  int64          `xml:"indOrgReemb"`
	Cnpjoper     []string       `xml:"cnpjOper,omitempty"`
	Regans       []string       `xml:"regANS,omitempty"`
	Detreembtit  []Detreembtit  `xml:"detReembTit"`
	Inforeembdep []Inforeembdep `xml:"infoReembDep"`
}

type Detreembtit struct {
	Tpinsc      int64    `xml:"tpInsc"`
	Nrinsc      string   `xml:"nrInsc"`
	Vlrreemb    []string `xml:"vlrReemb,omitempty"`
	Vlrreembant []string `xml:"vlrReembAnt,omitempty"`
}

type Inforeembdep struct {
	Cpfbenef string `xml:"cpfBenef"`
}

type S1260 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evtcomprod struct {
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Infocomprod   Infocomprod   `xml:"infoComProd"`
}

type Infocomprod struct {
	Ideestabel Ideestabel `xml:"ideEstabel"`
}

type Ideestabel struct {
	Nrinscestabrural string     `xml:"nrInscEstabRural"`
	Tpcomerc         []Tpcomerc `xml:"tpComerc"`
}

type Tpcomerc struct {
	Indcomerc   int64         `xml:"indComerc"`
	Vrtotcom    float64       `xml:"vrTotCom"`
	Ideadquir   []Ideadquir   `xml:"ideAdquir"`
	Infoprocjud []Infoprocjud `xml:"infoProcJud"`
}

type Ideadquir struct {
	Tpinsc   int64   `xml:"tpInsc"`
	Nrinsc   string  `xml:"nrInsc"`
	Vrcomerc float64 `xml:"vrComerc"`
	Nfs      []Nfs   `xml:"nfs"`
}

type Nfs struct {
	Serie       []string `xml:"serie,omitempty"`
	Nrdocto     string   `xml:"nrDocto"`
	Dtemisnf    string   `xml:"dtEmisNF"`
	Vlrbruto    float64  `xml:"vlrBruto"`
	Vrcpdescpr  float64  `xml:"vrCPDescPR"`
	Vrratdescpr float64  `xml:"vrRatDescPR"`
	Vrsenardesc float64  `xml:"vrSenarDesc"`
}

type Infoprocjud struct {
	Tpproc      int64    `xml:"tpProc"`
	Nrproc      string   `xml:"nrProc"`
	Codsusp     int64    `xml:"codSusp"`
	Vrcpsusp    []string `xml:"vrCPSusp,omitempty"`
	Vrratsusp   []string `xml:"vrRatSusp,omitempty"`
	Vrsenarsusp []string `xml:"vrSenarSusp,omitempty"`
	Dtsent      string   `xml:"dtSent"`
	Ufvara      string   `xml:"ufVara"`
	Codmunic    int64    `xml:"codMunic"`
	Idvara      int64    `xml:"idVara"`
}

type S1270 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evtcontratavnp struct {
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Remunavnp     []Remunavnp   `xml:"remunAvNP"`
}

type Remunavnp struct {
	Tpinsc     int64   `xml:"tpInsc"`
	Nrinsc     string  `xml:"nrInsc"`
	Codlotacao string  `xml:"codLotacao"`
	Vrbccp00   float64 `xml:"vrBcCp00"`
	Vrbccp15   float64 `xml:"vrBcCp15"`
	Vrbccp20   float64 `xml:"vrBcCp20"`
	Vrbccp25   float64 `xml:"vrBcCp25"`
	Vrbccp13   float64 `xml:"vrBcCp13"`
	Vrbcfgts   float64 `xml:"vrBcFgts"`
	Vrdesccp   float64 `xml:"vrDescCP"`
}

type S1280 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evtinfocomplper struct {
	Id                  string                `xml:"Id"`
	Ideevento           Ideevento             `xml:"ideEvento"`
	Ideempregador       Ideempregador         `xml:"ideEmpregador"`
	Infosubstpatr       []Infosubstpatr       `xml:"infoSubstPatr,omitempty"`
	Infosubstpatropport []Infosubstpatropport `xml:"infoSubstPatrOpPort"`
	Infoativconcom      []Infoativconcom      `xml:"infoAtivConcom,omitempty"`
	Infoperctransf11096 []Infoperctransf11096 `xml:"infoPercTransf11096,omitempty"`
}

type Infosubstpatr struct {
	Indsubstpatr   int64   `xml:"indSubstPatr"`
	Percredcontrib float64 `xml:"percRedContrib"`
}

type Infosubstpatropport struct {
	Codlotacao      string `xml:"codLotacao"`
	Cnpjopportuario string `xml:"cnpjOpPortuario"`
}

type Infoativconcom struct {
	Fatormes float64 `xml:"fatorMes"`
	Fator13  float64 `xml:"fator13"`
}

type Infoperctransf11096 struct {
	Perctransf int64 `xml:"percTransf"`
}

type S1298 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evtreabreevper struct {
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
}

type S1299 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evtfechaevper struct {
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Infofech      Infofech      `xml:"infoFech"`
}

type Infofech struct {
	Evtremun        string   `xml:"evtRemun"`
	Evtpgtos        string   `xml:"evtPgtos"`
	Evtcomprod      string   `xml:"evtComProd"`
	Evtcontratavnp  string   `xml:"evtContratAvNP"`
	Evtinfocomplper string   `xml:"evtInfoComplPer"`
	Indexcapur1250  []string `xml:"indExcApur1250,omitempty"`
	Transdctfweb    []string `xml:"transDCTFWeb,omitempty"`
	Naovalid        []string `xml:"naoValid,omitempty"`
}

type S2190 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evtadmprelim struct {
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Inforegprelim Inforegprelim `xml:"infoRegPrelim"`
}

type Inforegprelim struct {
	Cpftrab      string        `xml:"cpfTrab"`
	Dtnascto     string        `xml:"dtNascto"`
	Dtadm        string        `xml:"dtAdm"`
	Matricula    string        `xml:"matricula"`
	Codcateg     int64         `xml:"codCateg"`
	Natatividade []string      `xml:"natAtividade,omitempty"`
	Inforegctps  []Inforegctps `xml:"infoRegCTPS,omitempty"`
}

type Inforegctps struct {
	Cbocargo   string   `xml:"CBOCargo"`
	Vrsalfx    float64  `xml:"vrSalFx"`
	Undsalfixo int64    `xml:"undSalFixo"`
	Tpcontr    int64    `xml:"tpContr"`
	Dtterm     []string `xml:"dtTerm,omitempty"`
}

type S2200 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
	string  `xml:""`
}

type Evtadmissao struct {
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Trabalhador   Trabalhador   `xml:"trabalhador"`
	Vinculo       Vinculo       `xml:"vinculo"`
}

type Trabalhador struct {
	Cpftrab         string            `xml:"cpfTrab"`
	Nmtrab          string            `xml:"nmTrab"`
	Sexo            string            `xml:"sexo"`
	Racacor         int64             `xml:"racaCor"`
	Estciv          []string          `xml:"estCiv,omitempty"`
	Grauinstr       string            `xml:"grauInstr"`
	Nmsoc           []string          `xml:"nmSoc,omitempty"`
	Nascimento      Nascimento        `xml:"nascimento"`
	Endereco        []string          `xml:"endereco,omitempty"`
	Trabimig        []Trabimig        `xml:"trabImig,omitempty"`
	Infodeficiencia []Infodeficiencia `xml:"infoDeficiencia,omitempty"`
	Dependente      []Dependente      `xml:"dependente"`
	Contato         []Contato         `xml:"contato,omitempty"`
}

type Nascimento struct {
	Dtnascto   string `xml:"dtNascto"`
	Paisnascto string `xml:"paisNascto"`
	Paisnac    string `xml:"paisNac"`
}

type Endereco struct {
	Brasil   []Brasil   `xml:"brasil,omitempty"`
	Exterior []Exterior `xml:"exterior,omitempty"`
}

type Brasil struct {
	Tplograd    []string `xml:"tpLograd,omitempty"`
	Dsclograd   string   `xml:"dscLograd"`
	Nrlograd    string   `xml:"nrLograd"`
	Complemento []string `xml:"complemento,omitempty"`
	Bairro      []string `xml:"bairro,omitempty"`
	Cep         string   `xml:"cep"`
	Codmunic    int64    `xml:"codMunic"`
	Uf          string   `xml:"uf"`
}

type Exterior struct {
	Paisresid   string   `xml:"paisResid"`
	Dsclograd   string   `xml:"dscLograd"`
	Nrlograd    string   `xml:"nrLograd"`
	Complemento []string `xml:"complemento,omitempty"`
	Bairro      []string `xml:"bairro,omitempty"`
	Nmcid       string   `xml:"nmCid"`
	Codpostal   []string `xml:"codPostal,omitempty"`
}

type Trabimig struct {
	Tmpresid []string `xml:"tmpResid,omitempty"`
	Conding  int64    `xml:"condIng"`
}

type Infodeficiencia struct {
	Deffisica      string   `xml:"defFisica"`
	Defvisual      string   `xml:"defVisual"`
	Defauditiva    string   `xml:"defAuditiva"`
	Defmental      string   `xml:"defMental"`
	Defintelectual string   `xml:"defIntelectual"`
	Reabreadap     string   `xml:"reabReadap"`
	Observacao     []string `xml:"observacao,omitempty"`
}

type Dependente struct {
	Depsf     string   `xml:"depSF"`
	Inctrab   string   `xml:"incTrab"`
	Tpdep     []string `xml:"tpDep,omitempty"`
	Nmdep     string   `xml:"nmDep"`
	Dtnascto  string   `xml:"dtNascto"`
	Cpfdep    []string `xml:"cpfDep,omitempty"`
	Sexodep   string   `xml:"sexoDep"`
	Depirrf   string   `xml:"depIRRF"`
	Incfismen string   `xml:"incFisMen"`
	Descrdep  []string `xml:"descrDep,omitempty"`
}

type Contato struct {
	Foneprinc  []string `xml:"fonePrinc,omitempty"`
	Emailprinc []string `xml:"emailPrinc,omitempty"`
}

type Vinculo struct {
	Matricula      string           `xml:"matricula"`
	Tpregtrab      int64            `xml:"tpRegTrab"`
	Cadini         string           `xml:"cadIni"`
	Sucessaovinc   []Sucessaovinc   `xml:"sucessaoVinc,omitempty"`
	Transfdom      []Transfdom      `xml:"transfDom,omitempty"`
	Mudancacpf     []Mudancacpf     `xml:"mudancaCPF,omitempty"`
	Afastamento    []Afastamento    `xml:"afastamento,omitempty"`
	Desligamento   []Desligamento   `xml:"desligamento,omitempty"`
	Cessao         []Cessao         `xml:"cessao,omitempty"`
	Tpregprev      int64            `xml:"tpRegPrev"`
	Inforegimetrab []Inforegimetrab `xml:"infoRegimeTrab,omitempty"`
	Infocontrato   Infocontrato     `xml:"infoContrato"`
}

type Inforegimetrab struct {
	Infoceletista   []Infoceletista   `xml:"infoCeletista,omitempty"`
	Infoestatutario []Infoestatutario `xml:"infoEstatutario,omitempty"`
}

type Infoceletista struct {
	Dtadm             string           `xml:"dtAdm"`
	Tpadmissao        int64            `xml:"tpAdmissao"`
	Indadmissao       int64            `xml:"indAdmissao"`
	Nrproctrab        []string         `xml:"nrProcTrab,omitempty"`
	Matanotjud        []string         `xml:"matAnotJud,omitempty"`
	Fgts              []Fgts           `xml:"FGTS,omitempty"`
	Tpregjor          int64            `xml:"tpRegJor"`
	Natatividade      int64            `xml:"natAtividade"`
	Dtbase            []string         `xml:"dtBase,omitempty"`
	Cnpjsindcategprof string           `xml:"cnpjSindCategProf"`
	Trabtemporario    []Trabtemporario `xml:"trabTemporario,omitempty"`
	Aprend            []Aprend         `xml:"aprend,omitempty"`
}

type Fgts struct {
	Dtopcfgts string `xml:"dtOpcFGTS"`
}

type Trabtemporario struct {
	Hipleg             int64                `xml:"hipLeg"`
	Justcontr          string               `xml:"justContr"`
	Ideestabvinc       Ideestabvinc         `xml:"ideEstabVinc"`
	Idetrabsubstituido []Idetrabsubstituido `xml:"ideTrabSubstituido"`
	Justprorr          string               `xml:"justProrr"`
}

type Ideestabvinc struct {
	Tpinsc int64  `xml:"tpInsc"`
	Nrinsc string `xml:"nrInsc"`
}

type Idetrabsubstituido struct {
	Cpftrabsubst string `xml:"cpfTrabSubst"`
}

type Aprend struct {
	Indaprend   int64    `xml:"indAprend"`
	Cnpjentqual []string `xml:"cnpjEntQual,omitempty"`
	Tpinsc      []string `xml:"tpInsc,omitempty"`
	Nrinsc      []string `xml:"nrInsc,omitempty"`
	Cnpjprat    []string `xml:"cnpjPrat,omitempty"`
}

type Infoestatutario struct {
	Tpprov       int64    `xml:"tpProv"`
	Dtexercicio  string   `xml:"dtExercicio"`
	Dtiniabono   []string `xml:"dtIniAbono,omitempty"`
	Tpplanrp     int64    `xml:"tpPlanRP"`
	Indtetorgps  string   `xml:"indTetoRGPS"`
	Indabonoperm string   `xml:"indAbonoPerm"`
}

type Infocontrato struct {
	Dtingrcargo    []string         `xml:"dtIngrCargo,omitempty"`
	Nmcargo        []string         `xml:"nmCargo,omitempty"`
	Cbocargo       []string         `xml:"CBOCargo,omitempty"`
	Nmfuncao       []string         `xml:"nmFuncao,omitempty"`
	Cbofuncao      []string         `xml:"CBOFuncao,omitempty"`
	Acumcargo      []string         `xml:"acumCargo,omitempty"`
	Codcateg       int64            `xml:"codCateg"`
	Remuneracao    []Remuneracao    `xml:"remuneracao,omitempty"`
	Duracao        []Duracao        `xml:"duracao,omitempty"`
	Localtrabalho  Localtrabalho    `xml:"localTrabalho"`
	Horcontratual  []Horcontratual  `xml:"horContratual,omitempty"`
	Alvarajudicial []Alvarajudicial `xml:"alvaraJudicial,omitempty"`
	Observacoes    []Observacoes    `xml:"observacoes"`
	Treicap        []Treicap        `xml:"treiCap"`
}

type Remuneracao struct {
	Dtremun    string   `xml:"dtRemun"`
	Vrsalfx    float64  `xml:"vrSalFx"`
	Undsalfixo int64    `xml:"undSalFixo"`
	Dscsalvar  []string `xml:"dscSalVar,omitempty"`
}

type Duracao struct {
	Tpcontr   int64    `xml:"tpContr"`
	Dtterm    []string `xml:"dtTerm,omitempty"`
	Clauassec []string `xml:"clauAssec,omitempty"`
	Objdet    []string `xml:"objDet,omitempty"`
}

type Localtrabalho struct {
	Localtrabgeral []Localtrabgeral `xml:"localTrabGeral,omitempty"`
	Localtempdom   []Localtempdom   `xml:"localTempDom,omitempty"`
}

type Localtrabgeral struct {
	Tpinsc   int64    `xml:"tpInsc"`
	Nrinsc   string   `xml:"nrInsc"`
	Desccomp []string `xml:"descComp,omitempty"`
}

type Horcontratual struct {
	Qtdhrssem  []string `xml:"qtdHrsSem,omitempty"`
	Tpjornada  int64    `xml:"tpJornada"`
	Tmpparc    int64    `xml:"tmpParc"`
	Hornoturno []string `xml:"horNoturno,omitempty"`
	Dscjorn    string   `xml:"dscJorn"`
}

type Alvarajudicial struct {
	Nrprocjud string `xml:"nrProcJud"`
}

type Observacoes struct {
	Observacao string `xml:"observacao"`
}

type Treicap struct {
	Codtreicap int64 `xml:"codTreiCap"`
}

type Transfdom struct {
	Cpfsubstituido string   `xml:"cpfSubstituido"`
	Matricant      []string `xml:"matricAnt,omitempty"`
	Dttransf       string   `xml:"dtTransf"`
}

type Mudancacpf struct {
	Cpfant         string   `xml:"cpfAnt"`
	Nrbeneficioant string   `xml:"nrBeneficioAnt"`
	Dtaltcpf       string   `xml:"dtAltCPF"`
	Observacao     []string `xml:"observacao,omitempty"`
}

type Afastamento struct {
	Dtiniafast  string `xml:"dtIniAfast"`
	Codmotafast string `xml:"codMotAfast"`
}

type Desligamento struct {
	Mtvdeslig    string   `xml:"mtvDeslig"`
	Dtdeslig     string   `xml:"dtDeslig"`
	Dtprojfimapi []string `xml:"dtProjFimAPI,omitempty"`
}

type Cessao struct {
	Dtinicessao string `xml:"dtIniCessao"`
}

type S2205 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evtaltcadastral struct {
	Id             string         `xml:"Id"`
	Ideevento      Ideevento      `xml:"ideEvento"`
	Ideempregador  Ideempregador  `xml:"ideEmpregador"`
	Idetrabalhador Idetrabalhador `xml:"ideTrabalhador"`
	Alteracao      Alteracao      `xml:"alteracao"`
}

type Dadostrabalhador struct {
	Nmtrab          string            `xml:"nmTrab"`
	Sexo            string            `xml:"sexo"`
	Racacor         int64             `xml:"racaCor"`
	Estciv          []string          `xml:"estCiv,omitempty"`
	Grauinstr       string            `xml:"grauInstr"`
	Nmsoc           []string          `xml:"nmSoc,omitempty"`
	Paisnac         string            `xml:"paisNac"`
	Endereco        []string          `xml:"endereco,omitempty"`
	Trabimig        []Trabimig        `xml:"trabImig,omitempty"`
	Infodeficiencia []Infodeficiencia `xml:"infoDeficiencia,omitempty"`
	Dependente      []Dependente      `xml:"dependente"`
	Contato         []Contato         `xml:"contato,omitempty"`
}

type S2206 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evtaltcontratual struct {
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Idevinculo    Idevinculo    `xml:"ideVinculo"`
	Altcontratual Altcontratual `xml:"altContratual"`
}

type Idevinculo struct {
	Codcateg  []string `xml:"codCateg,omitempty"`
	Cpftrab   string   `xml:"cpfTrab"`
	Matricula string   `xml:"matricula"`
}

type Altcontratual struct {
	Dtalteracao string   `xml:"dtAlteracao"`
	Dtef        []string `xml:"dtEf,omitempty"`
	Dscalt      []string `xml:"dscAlt,omitempty"`
	Vinculo     Vinculo  `xml:"vinculo"`
}

type Localtempdom struct {
	Tplograd    []string `xml:"tpLograd,omitempty"`
	Dsclograd   string   `xml:"dscLograd"`
	Nrlograd    string   `xml:"nrLograd"`
	Complemento []string `xml:"complemento,omitempty"`
	Bairro      []string `xml:"bairro,omitempty"`
	Cep         string   `xml:"cep"`
	Codmunic    int64    `xml:"codMunic"`
	Uf          string   `xml:"uf"`
}

type S2210 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evtcat struct {
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Idevinculo    Idevinculo    `xml:"ideVinculo"`
	Cat           Cat           `xml:"cat"`
}

type Cat struct {
	Dtacid           string         `xml:"dtAcid"`
	Tpacid           int64          `xml:"tpAcid"`
	Hracid           []string       `xml:"hrAcid,omitempty"`
	Hrstrabantesacid []string       `xml:"hrsTrabAntesAcid,omitempty"`
	Tpcat            int64          `xml:"tpCat"`
	Indcatobito      string         `xml:"indCatObito"`
	Dtobito          []string       `xml:"dtObito,omitempty"`
	Indcomunpolicia  string         `xml:"indComunPolicia"`
	Codsitgeradora   int64          `xml:"codSitGeradora"`
	Iniciatcat       int64          `xml:"iniciatCAT"`
	Obscat           []string       `xml:"obsCAT,omitempty"`
	Ultdiatrab       []string       `xml:"ultDiaTrab,omitempty"`
	Houveafast       []string       `xml:"houveAfast,omitempty"`
	Localacidente    Localacidente  `xml:"localAcidente"`
	Parteatingida    Parteatingida  `xml:"parteAtingida"`
	Agentecausador   Agentecausador `xml:"agenteCausador"`
	Atestado         Atestado       `xml:"atestado"`
	Catorigem        []Catorigem    `xml:"catOrigem,omitempty"`
}

type Localacidente struct {
	Tplocal      int64          `xml:"tpLocal"`
	Dsclocal     []string       `xml:"dscLocal,omitempty"`
	Tplograd     []string       `xml:"tpLograd,omitempty"`
	Dsclograd    string         `xml:"dscLograd"`
	Nrlograd     string         `xml:"nrLograd"`
	Complemento  []string       `xml:"complemento,omitempty"`
	Bairro       []string       `xml:"bairro,omitempty"`
	Cep          []string       `xml:"cep,omitempty"`
	Codmunic     []string       `xml:"codMunic,omitempty"`
	Uf           []string       `xml:"uf,omitempty"`
	Pais         []string       `xml:"pais,omitempty"`
	Codpostal    []string       `xml:"codPostal,omitempty"`
	Idelocalacid []Idelocalacid `xml:"ideLocalAcid,omitempty"`
}

type Idelocalacid struct {
	Tpinsc int64  `xml:"tpInsc"`
	Nrinsc string `xml:"nrInsc"`
}

type Parteatingida struct {
	Codparteating int64 `xml:"codParteAting"`
	Lateralidade  int64 `xml:"lateralidade"`
}

type Agentecausador struct {
	Codagntcausador int64 `xml:"codAgntCausador"`
}

type Atestado struct {
	Dtatendimento string   `xml:"dtAtendimento"`
	Hratendimento string   `xml:"hrAtendimento"`
	Indinternacao string   `xml:"indInternacao"`
	Durtrat       int64    `xml:"durTrat"`
	Indafast      string   `xml:"indAfast"`
	Dsclesao      int64    `xml:"dscLesao"`
	Dsccomplesao  []string `xml:"dscCompLesao,omitempty"`
	Diagprovavel  []string `xml:"diagProvavel,omitempty"`
	Codcid        string   `xml:"codCID"`
	Observacao    []string `xml:"observacao,omitempty"`
	Emitente      Emitente `xml:"emitente"`
}

type Emitente struct {
	Nmemit string   `xml:"nmEmit"`
	Ideoc  int64    `xml:"ideOC"`
	Nroc   string   `xml:"nrOC"`
	Ufoc   []string `xml:"ufOC,omitempty"`
}

type Catorigem struct {
	Nrreccatorig string `xml:"nrRecCatOrig"`
}

type S2220 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evtmonit struct {
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Idevinculo    Idevinculo    `xml:"ideVinculo"`
	Exmedocup     Exmedocup     `xml:"exMedOcup"`
}

type Exmedocup struct {
	Tpexameocup int64       `xml:"tpExameOcup"`
	Aso         Aso         `xml:"aso"`
	Respmonit   []Respmonit `xml:"respMonit,omitempty"`
}

type Aso struct {
	Dtaso  string   `xml:"dtAso"`
	Resaso []string `xml:"resAso,omitempty"`
	Exame  []Exame  `xml:"exame"`
	Medico Medico   `xml:"medico"`
}

type Exame struct {
	Dtexm         string   `xml:"dtExm"`
	Procrealizado int64    `xml:"procRealizado"`
	Obsproc       []string `xml:"obsProc,omitempty"`
	Ordexame      []string `xml:"ordExame,omitempty"`
	Indresult     []string `xml:"indResult,omitempty"`
}

type Medico struct {
	Nmmed string   `xml:"nmMed"`
	Nrcrm []string `xml:"nrCRM,omitempty"`
	Ufcrm []string `xml:"ufCRM,omitempty"`
}

type Respmonit struct {
	Cpfresp []string `xml:"cpfResp,omitempty"`
	Nmresp  string   `xml:"nmResp"`
	Nrcrm   string   `xml:"nrCRM"`
	Ufcrm   string   `xml:"ufCRM"`
}

type S2221 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evttoxic struct {
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Idevinculo    Idevinculo    `xml:"ideVinculo"`
	Toxicologico  Toxicologico  `xml:"toxicologico"`
}

type Toxicologico struct {
	Dtexame     string   `xml:"dtExame"`
	Cnpjlab     string   `xml:"cnpjLab"`
	Codseqexame string   `xml:"codSeqExame"`
	Nmmed       string   `xml:"nmMed"`
	Nrcrm       []string `xml:"nrCRM,omitempty"`
	Ufcrm       []string `xml:"ufCRM,omitempty"`
}

type S2230 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evtafasttemp struct {
	Id              string          `xml:"Id"`
	Ideevento       Ideevento       `xml:"ideEvento"`
	Ideempregador   Ideempregador   `xml:"ideEmpregador"`
	Idevinculo      Idevinculo      `xml:"ideVinculo"`
	Infoafastamento Infoafastamento `xml:"infoAfastamento"`
}

type Infoafastamento struct {
	Iniafastamento []Iniafastamento `xml:"iniAfastamento,omitempty"`
	Inforetif      []Inforetif      `xml:"infoRetif,omitempty"`
	Fimafastamento []Fimafastamento `xml:"fimAfastamento,omitempty"`
}

type Iniafastamento struct {
	Dtiniafast     string         `xml:"dtIniAfast"`
	Codmotafast    string         `xml:"codMotAfast"`
	Infomesmomtv   []string       `xml:"infoMesmoMtv,omitempty"`
	Tpacidtransito []string       `xml:"tpAcidTransito,omitempty"`
	Observacao     []string       `xml:"observacao,omitempty"`
	Peraquis       []Peraquis     `xml:"perAquis,omitempty"`
	Infocessao     []Infocessao   `xml:"infoCessao,omitempty"`
	Infomandsind   []Infomandsind `xml:"infoMandSind,omitempty"`
	Infomandelet   []Infomandelet `xml:"infoMandElet,omitempty"`
}

type Peraquis struct {
	Dtinicio string   `xml:"dtInicio"`
	Dtfim    []string `xml:"dtFim,omitempty"`
}

type Infocessao struct {
	Cnpjcess  string      `xml:"cnpjCess"`
	Infonus   int64       `xml:"infOnus"`
	Inicessao []Inicessao `xml:"iniCessao,omitempty"`
	Fimcessao []Fimcessao `xml:"fimCessao,omitempty"`
}

type Infomandsind struct {
	Cnpjsind     string `xml:"cnpjSind"`
	Infonusremun int64  `xml:"infOnusRemun"`
}

type Infomandelet struct {
	Cnpjmandelet  string   `xml:"cnpjMandElet"`
	Categorig     int64    `xml:"categOrig"`
	Cnpjorig      string   `xml:"cnpjOrig"`
	Matricorig    string   `xml:"matricOrig"`
	Dtexercorig   string   `xml:"dtExercOrig"`
	Tpregtrab     int64    `xml:"tpRegTrab"`
	Indremuncargo []string `xml:"indRemunCargo,omitempty"`
	Tpregprev     int64    `xml:"tpRegPrev"`
}

type Inforetif struct {
	Origretif int64    `xml:"origRetif"`
	Tpproc    []string `xml:"tpProc,omitempty"`
	Nrproc    []string `xml:"nrProc,omitempty"`
}

type Fimafastamento struct {
	Dttermafast string `xml:"dtTermAfast"`
}

type S2231 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evtcessao struct {
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Idevinculo    Idevinculo    `xml:"ideVinculo"`
	Infocessao    string        `xml:"infoCessao"`
}

type Inicessao struct {
	Dtinicessao string `xml:"dtIniCessao"`
	Cnpjcess    string `xml:"cnpjCess"`
	Respremun   string `xml:"respRemun"`
}

type Fimcessao struct {
	Dttermcessao string `xml:"dtTermCessao"`
}

type S2240 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evtexprisco struct {
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Idevinculo    Idevinculo    `xml:"ideVinculo"`
	Infoexprisco  Infoexprisco  `xml:"infoExpRisco"`
}

type Infoexprisco struct {
	Dtinicondicao string    `xml:"dtIniCondicao"`
	Dtfimcondicao []string  `xml:"dtFimCondicao,omitempty"`
	Infoamb       []Infoamb `xml:"infoAmb"`
	Infoativ      Infoativ  `xml:"infoAtiv"`
	Agnoc         []Agnoc   `xml:"agNoc"`
	Respreg       []Respreg `xml:"respReg"`
	Obs           []Obs     `xml:"obs,omitempty"`
}

type Infoamb struct {
	Localamb int64  `xml:"localAmb"`
	Dscsetor string `xml:"dscSetor"`
	Tpinsc   int64  `xml:"tpInsc"`
	Nrinsc   string `xml:"nrInsc"`
}

type Infoativ struct {
	Dscativdes string `xml:"dscAtivDes"`
}

type Agnoc struct {
	Codagnoc   string   `xml:"codAgNoc"`
	Dscagnoc   []string `xml:"dscAgNoc,omitempty"`
	Tpaval     []string `xml:"tpAval,omitempty"`
	Intconc    []string `xml:"intConc,omitempty"`
	Limtol     []string `xml:"limTol,omitempty"`
	Unmed      []string `xml:"unMed,omitempty"`
	Tecmedicao []string `xml:"tecMedicao,omitempty"`
	Nrprocjud  []string `xml:"nrProcJud,omitempty"`
	Epcepi     []Epcepi `xml:"epcEpi,omitempty"`
}

type Epcepi struct {
	Utilizepc int64      `xml:"utilizEPC"`
	Eficepc   []string   `xml:"eficEpc,omitempty"`
	Utilizepi int64      `xml:"utilizEPI"`
	Eficepi   []string   `xml:"eficEpi,omitempty"`
	Epi       []Epi      `xml:"epi"`
	Epicompl  []Epicompl `xml:"epiCompl,omitempty"`
}

type Epi struct {
	Docaval string `xml:"docAval"`
}

type Epicompl struct {
	Medprotecao   string `xml:"medProtecao"`
	Condfuncto    string `xml:"condFuncto"`
	Usoinint      string `xml:"usoInint"`
	Przvalid      string `xml:"przValid"`
	Periodictroca string `xml:"periodicTroca"`
	Higienizacao  string `xml:"higienizacao"`
}

type Respreg struct {
	Cpfresp string   `xml:"cpfResp"`
	Ideoc   []string `xml:"ideOC,omitempty"`
	Dscoc   []string `xml:"dscOC,omitempty"`
	Nroc    []string `xml:"nrOC,omitempty"`
	Ufoc    []string `xml:"ufOC,omitempty"`
}

type Obs struct {
	Obscompl string `xml:"obsCompl"`
}

type S2298 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evtreintegr struct {
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Idevinculo    Idevinculo    `xml:"ideVinculo"`
	Inforeintegr  Inforeintegr  `xml:"infoReintegr"`
}

type Inforeintegr struct {
	Tpreint       int64    `xml:"tpReint"`
	Nrprocjud     []string `xml:"nrProcJud,omitempty"`
	Nrleianistia  []string `xml:"nrLeiAnistia,omitempty"`
	Dtefetretorno string   `xml:"dtEfetRetorno"`
	Dtefeito      string   `xml:"dtEfeito"`
}

type S2299 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evtdeslig struct {
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Idevinculo    Idevinculo    `xml:"ideVinculo"`
	Infodeslig    Infodeslig    `xml:"infoDeslig"`
}

type Infodeslig struct {
	Dtavprv         []string          `xml:"dtAvPrv,omitempty"`
	Indpagtoapi     string            `xml:"indPagtoAPI"`
	Nrproctrab      []string          `xml:"nrProcTrab,omitempty"`
	Indpdv          []string          `xml:"indPDV,omitempty"`
	Infointerm      []Infointerm      `xml:"infoInterm"`
	Observacoes     []Observacoes     `xml:"observacoes"`
	Sucessaovinc    []Sucessaovinc    `xml:"sucessaoVinc,omitempty"`
	Transftit       []Transftit       `xml:"transfTit,omitempty"`
	Mudancacpf      []Mudancacpf      `xml:"mudancaCPF,omitempty"`
	Verbasresc      []Verbasresc      `xml:"verbasResc,omitempty"`
	Remunaposdeslig []Remunaposdeslig `xml:"remunAposDeslig,omitempty"`
	Consigfgts      []Consigfgts      `xml:"consigFGTS"`
	Dtdeslig        string            `xml:"dtDeslig"`
	Mtvdeslig       string            `xml:"mtvDeslig"`
	Dtprojfimapi    []string          `xml:"dtProjFimAPI,omitempty"`
	Pensalim        []string          `xml:"pensAlim,omitempty"`
	Percaliment     []string          `xml:"percAliment,omitempty"`
	Vralim          []string          `xml:"vrAlim,omitempty"`
}

type Transftit struct {
	Cpfsubstituto string `xml:"cpfSubstituto"`
	Dtnascto      string `xml:"dtNascto"`
}

type Verbasresc struct {
	Proccs      []Proccs      `xml:"procCS,omitempty"`
	Dmdev       []Dmdev       `xml:"dmDev"`
	Procjudtrab []Procjudtrab `xml:"procJudTrab"`
	Infomv      []Infomv      `xml:"infoMV,omitempty"`
}

type Detverbas struct {
	Codrubr    string      `xml:"codRubr"`
	Idetabrubr string      `xml:"ideTabRubr"`
	Qtdrubr    []string    `xml:"qtdRubr,omitempty"`
	Fatorrubr  []string    `xml:"fatorRubr,omitempty"`
	Vrrubr     float64     `xml:"vrRubr"`
	Indapurir  []string    `xml:"indApurIR,omitempty"`
	Descfolha  []Descfolha `xml:"descFolha,omitempty"`
}

type Infosimples struct {
	Indsimples int64 `xml:"indSimples"`
}

type Proccs struct {
	Nrprocjud string `xml:"nrProcJud"`
}

type Remunaposdeslig struct {
	Indremun   []string `xml:"indRemun,omitempty"`
	Dtfimremun string   `xml:"dtFimRemun"`
}

type Consigfgts struct {
	Insconsig string `xml:"insConsig"`
	Nrcontr   string `xml:"nrContr"`
}

type S2300 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evttsvinicio struct {
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Trabalhador   Trabalhador   `xml:"trabalhador"`
	Infotsvinicio Infotsvinicio `xml:"infoTSVInicio"`
}

type Infotsvinicio struct {
	Cadini             string               `xml:"cadIni"`
	Matricula          []string             `xml:"matricula,omitempty"`
	Codcateg           int64                `xml:"codCateg"`
	Dtinicio           string               `xml:"dtInicio"`
	Nrproctrab         []string             `xml:"nrProcTrab,omitempty"`
	Natatividade       []string             `xml:"natAtividade,omitempty"`
	Infocomplementares []Infocomplementares `xml:"infoComplementares,omitempty"`
	Mudancacpf         []Mudancacpf         `xml:"mudancaCPF,omitempty"`
	Afastamento        []Afastamento        `xml:"afastamento,omitempty"`
	Termino            []Termino            `xml:"termino,omitempty"`
}

type Infocomplementares struct {
	Fgts                  []Fgts                  `xml:"FGTS,omitempty"`
	Cargofuncao           []Cargofuncao           `xml:"cargoFuncao,omitempty"`
	Remuneracao           []Remuneracao           `xml:"remuneracao,omitempty"`
	Infodirigentesindical []Infodirigentesindical `xml:"infoDirigenteSindical,omitempty"`
	Infotrabcedido        []Infotrabcedido        `xml:"infoTrabCedido,omitempty"`
	Infomandelet          []Infomandelet          `xml:"infoMandElet,omitempty"`
	Infoestagiario        []Infoestagiario        `xml:"infoEstagiario,omitempty"`
	Localtrabgeral        []Localtrabgeral        `xml:"localTrabGeral,omitempty"`
}

type Cargofuncao struct {
	Nmcargo   []string `xml:"nmCargo,omitempty"`
	Cbocargo  []string `xml:"CBOCargo,omitempty"`
	Nmfuncao  []string `xml:"nmFuncao,omitempty"`
	Cbofuncao []string `xml:"CBOFuncao,omitempty"`
}

type Infodirigentesindical struct {
	Categorig  int64    `xml:"categOrig"`
	Tpinsc     []string `xml:"tpInsc,omitempty"`
	Nrinsc     []string `xml:"nrInsc,omitempty"`
	Dtadmorig  []string `xml:"dtAdmOrig,omitempty"`
	Matricorig []string `xml:"matricOrig,omitempty"`
	Tpregtrab  []string `xml:"tpRegTrab,omitempty"`
	Tpregprev  int64    `xml:"tpRegPrev"`
}

type Infotrabcedido struct {
	Categorig int64  `xml:"categOrig"`
	Cnpjcednt string `xml:"cnpjCednt"`
	Matricced string `xml:"matricCed"`
	Dtadmced  string `xml:"dtAdmCed"`
	Tpregtrab int64  `xml:"tpRegTrab"`
	Tpregprev int64  `xml:"tpRegPrev"`
}

type Infoestagiario struct {
	Natestagio        string              `xml:"natEstagio"`
	Nivestagio        []string            `xml:"nivEstagio,omitempty"`
	Areaatuacao       []string            `xml:"areaAtuacao,omitempty"`
	Nrapol            []string            `xml:"nrApol,omitempty"`
	Dtprevterm        string              `xml:"dtPrevTerm"`
	Instensino        Instensino          `xml:"instEnsino"`
	Ageintegracao     []Ageintegracao     `xml:"ageIntegracao,omitempty"`
	Supervisorestagio []Supervisorestagio `xml:"supervisorEstagio,omitempty"`
}

type Instensino struct {
	Cnpjinstensino []string `xml:"cnpjInstEnsino,omitempty"`
	Nmrazao        []string `xml:"nmRazao,omitempty"`
	Dsclograd      []string `xml:"dscLograd,omitempty"`
	Nrlograd       []string `xml:"nrLograd,omitempty"`
	Bairro         []string `xml:"bairro,omitempty"`
	Cep            []string `xml:"cep,omitempty"`
	Codmunic       []string `xml:"codMunic,omitempty"`
	Uf             []string `xml:"uf,omitempty"`
}

type Ageintegracao struct {
	Cnpjagntinteg string `xml:"cnpjAgntInteg"`
}

type Supervisorestagio struct {
	Cpfsupervisor string `xml:"cpfSupervisor"`
}

type Termino struct {
	Dtterm string `xml:"dtTerm"`
}

type S2306 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evttsvaltcontr struct {
	Id                string            `xml:"Id"`
	Ideevento         Ideevento         `xml:"ideEvento"`
	Ideempregador     Ideempregador     `xml:"ideEmpregador"`
	Idetrabsemvinculo Idetrabsemvinculo `xml:"ideTrabSemVinculo"`
	Infotsvalteracao  Infotsvalteracao  `xml:"infoTSVAlteracao"`
}

type Idetrabsemvinculo struct {
	Cpftrab   string   `xml:"cpfTrab"`
	Matricula []string `xml:"matricula,omitempty"`
	Codcateg  []string `xml:"codCateg,omitempty"`
}

type Infotsvalteracao struct {
	Dtalteracao        string               `xml:"dtAlteracao"`
	Natatividade       []string             `xml:"natAtividade,omitempty"`
	Infocomplementares []Infocomplementares `xml:"infoComplementares,omitempty"`
}

type S2399 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evttsvtermino struct {
	Id                string            `xml:"Id"`
	Ideevento         Ideevento         `xml:"ideEvento"`
	Ideempregador     Ideempregador     `xml:"ideEmpregador"`
	Idetrabsemvinculo Idetrabsemvinculo `xml:"ideTrabSemVinculo"`
	Infotsvtermino    Infotsvtermino    `xml:"infoTSVTermino"`
}

type Infotsvtermino struct {
	Dtterm        string          `xml:"dtTerm"`
	Mtvdesligtsv  []string        `xml:"mtvDesligTSV,omitempty"`
	Pensalim      []string        `xml:"pensAlim,omitempty"`
	Percaliment   []string        `xml:"percAliment,omitempty"`
	Vralim        []string        `xml:"vrAlim,omitempty"`
	Nrproctrab    []string        `xml:"nrProcTrab,omitempty"`
	Mudancacpf    []Mudancacpf    `xml:"mudancaCPF,omitempty"`
	Verbasresc    []Verbasresc    `xml:"verbasResc,omitempty"`
	Remunaposterm []Remunaposterm `xml:"remunAposTerm,omitempty"`
}

type Remunaposterm struct {
	Indremun   []string `xml:"indRemun,omitempty"`
	Dtfimremun string   `xml:"dtFimRemun"`
}

type S2400 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evtcdbenefin struct {
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Beneficiario  Beneficiario  `xml:"beneficiario"`
}

type Beneficiario struct {
	Nmbenefic   string       `xml:"nmBenefic"`
	Dtnascto    string       `xml:"dtNascto"`
	Dtinicio    string       `xml:"dtInicio"`
	Sexo        []string     `xml:"sexo,omitempty"`
	Racacor     int64        `xml:"racaCor"`
	Estciv      []string     `xml:"estCiv,omitempty"`
	Incfismen   string       `xml:"incFisMen"`
	Dtincfismen []string     `xml:"dtIncFisMen,omitempty"`
	Endereco    string       `xml:"endereco"`
	Dependente  []Dependente `xml:"dependente"`
	Cpfbenef    string       `xml:"cpfBenef"`
	Matricula   []string     `xml:"matricula,omitempty"`
	Cnpjorigem  []string     `xml:"cnpjOrigem,omitempty"`
}

type S2405 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evtcdbenefalt struct {
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Idebenef      Idebenef      `xml:"ideBenef"`
	Alteracao     Alteracao     `xml:"alteracao"`
}

type Dadosbenef struct {
	Nmbenefic  string       `xml:"nmBenefic"`
	Sexo       string       `xml:"sexo"`
	Racacor    int64        `xml:"racaCor"`
	Estciv     []string     `xml:"estCiv,omitempty"`
	Incfismen  string       `xml:"incFisMen"`
	Endereco   string       `xml:"endereco"`
	Dependente []Dependente `xml:"dependente"`
}

type S2410 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evtcdbenin struct {
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Beneficiario  Beneficiario  `xml:"beneficiario"`
	Infobeninicio Infobeninicio `xml:"infoBenInicio"`
}

type Infobeninicio struct {
	Cadini         string           `xml:"cadIni"`
	Indsitbenef    []string         `xml:"indSitBenef,omitempty"`
	Nrbeneficio    string           `xml:"nrBeneficio"`
	Dtinibeneficio string           `xml:"dtIniBeneficio"`
	Dtpublic       []string         `xml:"dtPublic,omitempty"`
	Dadosbeneficio Dadosbeneficio   `xml:"dadosBeneficio"`
	Sucessaobenef  []Sucessaobenef  `xml:"sucessaoBenef,omitempty"`
	Mudancacpf     []Mudancacpf     `xml:"mudancaCPF,omitempty"`
	Infobentermino []Infobentermino `xml:"infoBenTermino,omitempty"`
}

type Dadosbeneficio struct {
	Inddecjud    []string       `xml:"indDecJud,omitempty"`
	Infohomolog  []Infohomolog  `xml:"infoHomolog,omitempty"`
	Tpbeneficio  string         `xml:"tpBeneficio"`
	Tpplanrp     int64          `xml:"tpPlanRP"`
	Dsc          []string       `xml:"dsc,omitempty"`
	Indsuspensao string         `xml:"indSuspensao"`
	Infopenmorte []Infopenmorte `xml:"infoPenMorte,omitempty"`
	Suspensao    []Suspensao    `xml:"suspensao,omitempty"`
}

type Infopenmorte struct {
	Tppenmorte   int64          `xml:"tpPenMorte"`
	Instpenmorte []Instpenmorte `xml:"instPenMorte,omitempty"`
}

type Instpenmorte struct {
	Cpfinst      string   `xml:"cpfInst"`
	Dtinst       string   `xml:"dtInst"`
	Tpdepinst    string   `xml:"tpDepInst"`
	Descrdepinst []string `xml:"descrDepInst,omitempty"`
}

type Infohomolog struct {
	Sithomolog int64    `xml:"sitHomolog"`
	Dthomolog  []string `xml:"dtHomolog,omitempty"`
}

type Sucessaobenef struct {
	Cnpjorgaoant   string   `xml:"cnpjOrgaoAnt"`
	Nrbeneficioant string   `xml:"nrBeneficioAnt"`
	Dttransf       string   `xml:"dtTransf"`
	Observacao     []string `xml:"observacao,omitempty"`
}

type Infobentermino struct {
	Dttermbeneficio string   `xml:"dtTermBeneficio"`
	Mtvtermino      string   `xml:"mtvTermino"`
	Cnpjorgaosuc    []string `xml:"cnpjOrgaoSuc,omitempty"`
	Novocpf         []string `xml:"novoCPF,omitempty"`
}

type S2416 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evtcdbenalt struct {
	Id               string           `xml:"Id"`
	Ideevento        Ideevento        `xml:"ideEvento"`
	Ideempregador    Ideempregador    `xml:"ideEmpregador"`
	Idebeneficio     Idebeneficio     `xml:"ideBeneficio"`
	Infobenalteracao Infobenalteracao `xml:"infoBenAlteracao"`
}

type Idebeneficio struct {
	Cpfbenef    string `xml:"cpfBenef"`
	Nrbeneficio string `xml:"nrBeneficio"`
}

type Infobenalteracao struct {
	Dtaltbeneficio string         `xml:"dtAltBeneficio"`
	Dadosbeneficio Dadosbeneficio `xml:"dadosBeneficio"`
}

type Suspensao struct {
	Mtvsuspensao string   `xml:"mtvSuspensao"`
	Dscsuspensao []string `xml:"dscSuspensao,omitempty"`
}

type S2418 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evtreativben struct {
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Idebeneficio  Idebeneficio  `xml:"ideBeneficio"`
	Inforeativ    Inforeativ    `xml:"infoReativ"`
}

type Inforeativ struct {
	Dtefetreativ string `xml:"dtEfetReativ"`
	Dtefeito     string `xml:"dtEfeito"`
}

type S2420 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evtcdbenterm struct {
	Id             string         `xml:"Id"`
	Ideevento      Ideevento      `xml:"ideEvento"`
	Ideempregador  Ideempregador  `xml:"ideEmpregador"`
	Idebeneficio   Idebeneficio   `xml:"ideBeneficio"`
	Infobentermino Infobentermino `xml:"infoBenTermino"`
}

type S2500 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evtproctrab struct {
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Infoprocesso  Infoprocesso  `xml:"infoProcesso"`
	Idetrab       Idetrab       `xml:"ideTrab"`
}

type Ideresp struct {
	Tpinsc       int64    `xml:"tpInsc"`
	Nrinsc       string   `xml:"nrInsc"`
	Dtadmrespdir []string `xml:"dtAdmRespDir,omitempty"`
	Matrespdir   []string `xml:"matRespDir,omitempty"`
}

type Dadoscompl struct {
	Infoprocjud []Infoprocjud `xml:"infoProcJud,omitempty"`
	Infoccp     []Infoccp     `xml:"infoCCP,omitempty"`
}

type Infoccp struct {
	Dtccp   string   `xml:"dtCCP"`
	Tpccp   int64    `xml:"tpCCP"`
	Cnpjccp []string `xml:"cnpjCCP,omitempty"`
}

type Idetrab struct {
	Nmtrab        []string        `xml:"nmTrab,omitempty"`
	Dtnascto      []string        `xml:"dtNascto,omitempty"`
	Ideseqtrab    []string        `xml:"ideSeqTrab,omitempty"`
	Infocontr     []Infocontr     `xml:"infoContr"`
	Cpftrab       string          `xml:"cpfTrab"`
	Calctrib      []Calctrib      `xml:"calcTrib"`
	Infocrirrf    []Infocrirrf    `xml:"infoCRIRRF"`
	Infoircomplem []Infoircomplem `xml:"infoIRComplem,omitempty"`
}

type Infocontr struct {
	Tpcontr      int64          `xml:"tpContr"`
	Indcontr     string         `xml:"indContr"`
	Dtadmorig    []string       `xml:"dtAdmOrig,omitempty"`
	Indreint     []string       `xml:"indReint,omitempty"`
	Indcateg     string         `xml:"indCateg"`
	Indnatativ   string         `xml:"indNatAtiv"`
	Indmotdeslig string         `xml:"indMotDeslig"`
	Matricula    []string       `xml:"matricula,omitempty"`
	Codcateg     []string       `xml:"codCateg,omitempty"`
	Dtinicio     []string       `xml:"dtInicio,omitempty"`
	Infocompl    []Infocompl    `xml:"infoCompl,omitempty"`
	Mudcategativ []Mudcategativ `xml:"mudCategAtiv"`
	Uniccontr    []Uniccontr    `xml:"unicContr"`
	Ideestab     Ideestab       `xml:"ideEstab"`
}

type Infocompl struct {
	Codcbo        []string        `xml:"codCBO,omitempty"`
	Natatividade  []string        `xml:"natAtividade,omitempty"`
	Remuneracao   []Remuneracao   `xml:"remuneracao"`
	Infovinc      []Infovinc      `xml:"infoVinc,omitempty"`
	Infoterm      []Infoterm      `xml:"infoTerm,omitempty"`
	Sucessaovinc  []Sucessaovinc  `xml:"sucessaoVinc,omitempty"`
	Infointerm    []Infointerm    `xml:"infoInterm"`
	Infocomplcont []Infocomplcont `xml:"infoComplCont"`
}

type Infovinc struct {
	Tpregtrab    int64          `xml:"tpRegTrab"`
	Tpregprev    int64          `xml:"tpRegPrev"`
	Dtadm        string         `xml:"dtAdm"`
	Tmpparc      []string       `xml:"tmpParc,omitempty"`
	Duracao      []Duracao      `xml:"duracao,omitempty"`
	Observacoes  []Observacoes  `xml:"observacoes"`
	Sucessaovinc []Sucessaovinc `xml:"sucessaoVinc,omitempty"`
	Infodeslig   []Infodeslig   `xml:"infoDeslig,omitempty"`
}

type Infoterm struct {
	Dtterm       string   `xml:"dtTerm"`
	Mtvdesligtsv []string `xml:"mtvDesligTSV,omitempty"`
}

type Mudcategativ struct {
	Codcateg       int64    `xml:"codCateg"`
	Natatividade   []string `xml:"natAtividade,omitempty"`
	Dtmudcategativ string   `xml:"dtMudCategAtiv"`
}

type Uniccontr struct {
	Matunic  []string `xml:"matUnic,omitempty"`
	Codcateg []string `xml:"codCateg,omitempty"`
	Dtinicio []string `xml:"dtInicio,omitempty"`
}

type Infovlr struct {
	Compini    string       `xml:"compIni"`
	Compfim    string       `xml:"compFim"`
	Indreperc  int64        `xml:"indReperc"`
	Indensd    []string     `xml:"indenSD,omitempty"`
	Indenabono []string     `xml:"indenAbono,omitempty"`
	Abono      []Abono      `xml:"abono"`
	Ideperiodo []Ideperiodo `xml:"idePeriodo"`
}

type Abono struct {
	Anobase int64 `xml:"anoBase"`
}

type Basecalculo struct {
	Vrbccpmensal float64        `xml:"vrBcCpMensal"`
	Vrbccp13     []string       `xml:"vrBcCp13,omitempty"`
	Infoagnocivo []Infoagnocivo `xml:"infoAgNocivo,omitempty"`
}

type Infofgts struct {
	Vrbcfgtsproctrab float64    `xml:"vrBcFGTSProcTrab"`
	Vrbcfgtssefip    []string   `xml:"vrBcFGTSSefip,omitempty"`
	Vrbcfgtsdecant   []string   `xml:"vrBcFGTSDecAnt,omitempty"`
	Dtvenc           []string   `xml:"dtVenc,omitempty"`
	Classtrib        []string   `xml:"classTrib,omitempty"`
	Nrrecarqbase     string     `xml:"nrRecArqBase"`
	Indexistinfo     int64      `xml:"indExistInfo"`
	Ideestab         []Ideestab `xml:"ideEstab"`
}

type Basemudcateg struct {
	Codcateg  int64   `xml:"codCateg"`
	Vrbccprev float64 `xml:"vrBcCPrev"`
}

type S2501 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evtcontproc struct {
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Ideproc       Ideproc       `xml:"ideProc"`
	Idetrab       []Idetrab     `xml:"ideTrab"`
}

type Ideproc struct {
	Ideseqproc   []string       `xml:"ideSeqProc,omitempty"`
	Obs          []string       `xml:"obs,omitempty"`
	Perapurpgto  string         `xml:"perApurPgto"`
	Perapur      string         `xml:"perApur"`
	Infotributos []Infotributos `xml:"infoTributos"`
	Infocrirrf   []Infocrirrf   `xml:"infoCRIRRF"`
	Origem       int64          `xml:"origem"`
	Nrproctrab   string         `xml:"nrProcTrab"`
}

type Calctrib struct {
	Perref        string          `xml:"perRef"`
	Vrbccpmensal  float64         `xml:"vrBcCpMensal"`
	Vrbccp13      float64         `xml:"vrBcCp13"`
	Infocrcontrib []Infocrcontrib `xml:"infoCRContrib"`
}

type Infocrcontrib struct {
	Vrcrsusp []string `xml:"vrCRSusp,omitempty"`
	Tpcr     int64    `xml:"tpCR"`
	Vrcr     float64  `xml:"vrCR"`
}

type Infocrirrf struct {
	Vrcr13      []string      `xml:"vrCR13,omitempty"`
	Infoir      []Infoir      `xml:"infoIR,omitempty"`
	Inforra     []Inforra     `xml:"infoRRA,omitempty"`
	Deddepen    []Deddepen    `xml:"dedDepen"`
	Penalim     []Penalim     `xml:"penAlim"`
	Infoprocret []Infoprocret `xml:"infoProcRet"`
	Tpcr        int64         `xml:"tpCR"`
	Vrcr        float64       `xml:"vrCR"`
}

type Infoir struct {
	Vrrendtrib        []string         `xml:"vrRendTrib,omitempty"`
	Vrrendtrib13      []string         `xml:"vrRendTrib13,omitempty"`
	Vrrendmolegrave   []string         `xml:"vrRendMoleGrave,omitempty"`
	Vrrendmolegrave13 []string         `xml:"vrRendMoleGrave13,omitempty"`
	Vrrendisen65      []string         `xml:"vrRendIsen65,omitempty"`
	Vrrendisen65dec   []string         `xml:"vrRendIsen65Dec,omitempty"`
	Vrjurosmora       []string         `xml:"vrJurosMora,omitempty"`
	Vrjurosmora13     []string         `xml:"vrJurosMora13,omitempty"`
	Vrrendisenntrib   []string         `xml:"vrRendIsenNTrib,omitempty"`
	Descisenntrib     []string         `xml:"descIsenNTrib,omitempty"`
	Vrprevoficial     []string         `xml:"vrPrevOficial,omitempty"`
	Vrprevoficial13   []string         `xml:"vrPrevOficial13,omitempty"`
	Rendisen0561      []Rendisen0561   `xml:"rendIsen0561,omitempty"`
	Tpinfoir          int64            `xml:"tpInfoIR"`
	Valor             float64          `xml:"valor"`
	Descrendimento    []string         `xml:"descRendimento,omitempty"`
	Infoprocjudrub    []Infoprocjudrub `xml:"infoProcJudRub"`
}

type Rendisen0561 struct {
	Vlrdiarias        []string `xml:"vlrDiarias,omitempty"`
	Vlrajudacusto     []string `xml:"vlrAjudaCusto,omitempty"`
	Vlrindrescontrato []string `xml:"vlrIndResContrato,omitempty"`
	Vlrabonopec       []string `xml:"vlrAbonoPec,omitempty"`
	Vlrauxmoradia     []string `xml:"vlrAuxMoradia,omitempty"`
}

type S2555 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evtconsolidcontproc struct {
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Ideproc       Ideproc       `xml:"ideProc"`
}

type S3000 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evtexclusao struct {
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Infoexclusao  Infoexclusao  `xml:"infoExclusao"`
}

type Infoexclusao struct {
	Idetrabalhador []Idetrabalhador `xml:"ideTrabalhador,omitempty"`
	Idefolhapagto  []Idefolhapagto  `xml:"ideFolhaPagto,omitempty"`
	Tpevento       string           `xml:"tpEvento"`
	Nrrecevt       string           `xml:"nrRecEvt"`
	Ideproctrab    Ideproctrab      `xml:"ideProcTrab"`
}

type Idefolhapagto struct {
	Indapuracao []string `xml:"indApuracao,omitempty"`
	Perapur     string   `xml:"perApur"`
}

type S3500 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evtexcproctrab struct {
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Infoexclusao  Infoexclusao  `xml:"infoExclusao"`
}

type Ideproctrab struct {
	Nrproctrab  string   `xml:"nrProcTrab"`
	Cpftrab     []string `xml:"cpfTrab,omitempty"`
	Perapurpgto []string `xml:"perApurPgto,omitempty"`
	Ideseqproc  []string `xml:"ideSeqProc,omitempty"`
}

type S5001 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evtbasestrab struct {
	Id             string         `xml:"Id"`
	Ideevento      Ideevento      `xml:"ideEvento"`
	Ideempregador  Ideempregador  `xml:"ideEmpregador"`
	Idetrabalhador Idetrabalhador `xml:"ideTrabalhador"`
	Infocpcalc     []Infocpcalc   `xml:"infoCpCalc"`
	Infocp         []Infocp       `xml:"infoCp,omitempty"`
	Infopispasep   []Infopispasep `xml:"infoPisPasep,omitempty"`
}

type Infocpcalc struct {
	Tpcr      int64   `xml:"tpCR"`
	Vrcpseg   float64 `xml:"vrCpSeg"`
	Vrdescseg float64 `xml:"vrDescSeg"`
}

type Infocp struct {
	Classtrib   string        `xml:"classTrib"`
	Ideestablot []Ideestablot `xml:"ideEstabLot"`
}

type Infocategincid struct {
	Matricula  []string     `xml:"matricula,omitempty"`
	Codcateg   int64        `xml:"codCateg"`
	Indsimples []string     `xml:"indSimples,omitempty"`
	Infobasecs []Infobasecs `xml:"infoBaseCS"`
	Calcterc   []Calcterc   `xml:"calcTerc"`
	Infoperref []Infoperref `xml:"infoPerRef"`
}

type Infobasecs struct {
	Ind13   int64   `xml:"ind13"`
	Tpvalor int64   `xml:"tpValor"`
	Valor   float64 `xml:"valor"`
}

type Calcterc struct {
	Tpcr        int64   `xml:"tpCR"`
	Vrcssegterc float64 `xml:"vrCsSegTerc"`
	Vrdescterc  float64 `xml:"vrDescTerc"`
}

type Infoperref struct {
	Perref        string          `xml:"perRef"`
	Ideadc        []Ideadc        `xml:"ideADC"`
	Detinfoperref []Detinfoperref `xml:"detInfoPerRef"`
}

type Detinfoperref struct {
	Ind13      int64   `xml:"ind13"`
	Tpvrperref int64   `xml:"tpVrPerRef"`
	Vrperref   float64 `xml:"vrPerRef"`
}

type Infopispasep struct {
	Ideestab []Ideestab `xml:"ideEstab"`
}

type Infocategpispasep struct {
	Matricula        []string           `xml:"matricula,omitempty"`
	Codcateg         int64              `xml:"codCateg"`
	Infobasepispasep []Infobasepispasep `xml:"infoBasePisPasep"`
}

type Infobasepispasep struct {
	Ind13           int64   `xml:"ind13"`
	Tpvalorpispasep int64   `xml:"tpValorPisPasep"`
	Valorpispasep   float64 `xml:"valorPisPasep"`
}

type S5002 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
	string  `xml:""`
}

type Evtirrfbenef struct {
	Id             string         `xml:"Id"`
	Ideevento      Ideevento      `xml:"ideEvento"`
	Ideempregador  Ideempregador  `xml:"ideEmpregador"`
	Idetrabalhador Idetrabalhador `xml:"ideTrabalhador"`
}

type Infoprocjudrub struct {
	Nrproc   string `xml:"nrProc"`
	Ufvara   string `xml:"ufVara"`
	Codmunic int64  `xml:"codMunic"`
	Idvara   int64  `xml:"idVara"`
}

type Totapurmen struct {
	Crmen              string   `xml:"CRMen"`
	Vlrrendtrib        float64  `xml:"vlrRendTrib"`
	Vlrrendtrib13      float64  `xml:"vlrRendTrib13"`
	Vlrprevoficial     float64  `xml:"vlrPrevOficial"`
	Vlrprevoficial13   float64  `xml:"vlrPrevOficial13"`
	Vlrcrmen           float64  `xml:"vlrCRMen"`
	Vlrcr13men         float64  `xml:"vlrCR13Men"`
	Vlrparcisenta65    float64  `xml:"vlrParcIsenta65"`
	Vlrparcisenta65dec float64  `xml:"vlrParcIsenta65Dec"`
	Vlrdiarias         float64  `xml:"vlrDiarias"`
	Vlrajudacusto      float64  `xml:"vlrAjudaCusto"`
	Vlrindrescontrato  float64  `xml:"vlrIndResContrato"`
	Vlrabonopec        float64  `xml:"vlrAbonoPec"`
	Vlrrendmolegrave   float64  `xml:"vlrRendMoleGrave"`
	Vlrrendmolegrave13 float64  `xml:"vlrRendMoleGrave13"`
	Vlrauxmoradia      float64  `xml:"vlrAuxMoradia"`
	Vlrbolsamedico     float64  `xml:"vlrBolsaMedico"`
	Vlrbolsamedico13   float64  `xml:"vlrBolsaMedico13"`
	Vlrjurosmora       float64  `xml:"vlrJurosMora"`
	Vlrisenoutros      float64  `xml:"vlrIsenOutros"`
	Descrendimento     []string `xml:"descRendimento,omitempty"`
}

type Totapurdia struct {
	Perapurdia   int64   `xml:"perApurDia"`
	Crdia        string  `xml:"CRDia"`
	Frmtribut    string  `xml:"frmTribut"`
	Paisresidext string  `xml:"paisResidExt"`
	Vlrpagodia   float64 `xml:"vlrPagoDia"`
	Vlrcrdia     float64 `xml:"vlrCRDia"`
}

type Totinfoir struct {
	Consolidapurmen []Consolidapurmen `xml:"consolidApurMen"`
}

type Consolidapurmen struct {
	Crmen              string   `xml:"CRMen"`
	Vlrrendtrib        float64  `xml:"vlrRendTrib"`
	Vlrrendtrib13      float64  `xml:"vlrRendTrib13"`
	Vlrprevoficial     float64  `xml:"vlrPrevOficial"`
	Vlrprevoficial13   float64  `xml:"vlrPrevOficial13"`
	Vlrcrmen           float64  `xml:"vlrCRMen"`
	Vlrcr13men         float64  `xml:"vlrCR13Men"`
	Vlrparcisenta65    float64  `xml:"vlrParcIsenta65"`
	Vlrparcisenta65dec float64  `xml:"vlrParcIsenta65Dec"`
	Vlrdiarias         float64  `xml:"vlrDiarias"`
	Vlrajudacusto      float64  `xml:"vlrAjudaCusto"`
	Vlrindrescontrato  float64  `xml:"vlrIndResContrato"`
	Vlrabonopec        float64  `xml:"vlrAbonoPec"`
	Vlrrendmolegrave   float64  `xml:"vlrRendMoleGrave"`
	Vlrrendmolegrave13 float64  `xml:"vlrRendMoleGrave13"`
	Vlrauxmoradia      float64  `xml:"vlrAuxMoradia"`
	Vlrbolsamedico     float64  `xml:"vlrBolsaMedico"`
	Vlrbolsamedico13   float64  `xml:"vlrBolsaMedico13"`
	Vlrjurosmora       float64  `xml:"vlrJurosMora"`
	Vlrisenoutros      float64  `xml:"vlrIsenOutros"`
	Descrendimento     []string `xml:"descRendimento,omitempty"`
}

type Idedep struct {
	Cpfdep   string   `xml:"cpfDep"`
	Depirrf  []string `xml:"depIRRF,omitempty"`
	Dtnascto []string `xml:"dtNascto,omitempty"`
	Nome     []string `xml:"nome,omitempty"`
	Tpdep    []string `xml:"tpDep,omitempty"`
	Descrdep []string `xml:"descrDep,omitempty"`
}

type S5003 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
	string  `xml:""`
}

type Evtbasesfgts struct {
	Id             string         `xml:"Id"`
	Ideevento      Ideevento      `xml:"ideEvento"`
	Ideempregador  Ideempregador  `xml:"ideEmpregador"`
	Idetrabalhador Idetrabalhador `xml:"ideTrabalhador"`
	Infofgts       Infofgts       `xml:"infoFGTS"`
}

type Infotrabfgts struct {
	Tpregtrab        []string         `xml:"tpRegTrab,omitempty"`
	Remunsuc         []string         `xml:"remunSuc,omitempty"`
	Dtdeslig         []string         `xml:"dtDeslig,omitempty"`
	Mtvdeslig        []string         `xml:"mtvDeslig,omitempty"`
	Dtterm           []string         `xml:"dtTerm,omitempty"`
	Mtvdesligtsv     []string         `xml:"mtvDesligTSV,omitempty"`
	Sucessaovinc     []Sucessaovinc   `xml:"sucessaoVinc,omitempty"`
	Infobasefgts     []Infobasefgts   `xml:"infoBaseFGTS,omitempty"`
	Proccs           []Proccs         `xml:"procCS,omitempty"`
	Econsignado      []Econsignado    `xml:"eConsignado"`
	Matricula        []string         `xml:"matricula,omitempty"`
	Codcateg         []string         `xml:"codCateg,omitempty"`
	Categorig        []string         `xml:"categOrig,omitempty"`
	Infofgtsproctrab Infofgtsproctrab `xml:"infoFGTSProcTrab"`
}

type Infobasefgts struct {
	Baseperapur     []Baseperapur     `xml:"basePerApur"`
	Infobaseperante []Infobaseperante `xml:"infoBasePerAntE"`
}

type Baseperapur struct {
	Remfgts     float64       `xml:"remFGTS"`
	Dpsfgts     []string      `xml:"dpsFGTS,omitempty"`
	Detrubrsusp []Detrubrsusp `xml:"detRubrSusp"`
	Tpvalor     int64         `xml:"tpValor"`
	Indincid    int64         `xml:"indIncid"`
	Basefgts    float64       `xml:"baseFGTS"`
	Vrfgts      []string      `xml:"vrFGTS,omitempty"`
	Notaft      []string      `xml:"notAFT,omitempty"`
	Natrubr     []string      `xml:"natRubr,omitempty"`
}

type Detrubrsusp struct {
	Codrubr         string            `xml:"codRubr"`
	Idetabrubr      string            `xml:"ideTabRubr"`
	Vrrubr          float64           `xml:"vrRubr"`
	Ideprocessofgts []Ideprocessofgts `xml:"ideProcessoFGTS"`
}

type Infobaseperante struct {
	Perref      string        `xml:"perRef"`
	Tpacconv    string        `xml:"tpAcConv"`
	Baseperante []Baseperante `xml:"basePerAntE"`
}

type Baseperante struct {
	Remfgtse    float64       `xml:"remFGTSE"`
	Dpsfgtse    []string      `xml:"dpsFGTSE,omitempty"`
	Detrubrsusp []Detrubrsusp `xml:"detRubrSusp"`
	Tpvalore    int64         `xml:"tpValorE"`
	Indincide   int64         `xml:"indIncidE"`
	Basefgtse   float64       `xml:"baseFGTSE"`
	Vrfgtse     []string      `xml:"vrFGTSE,omitempty"`
}

type Econsignado struct {
	Instfinanc    string  `xml:"instFinanc"`
	Nrcontrato    string  `xml:"nrContrato"`
	Vreconsignado float64 `xml:"vreConsignado"`
}

type S5011 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evtcs struct {
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Infocs        Infocs        `xml:"infoCS"`
}

type Infocs struct {
	Nrrecarqbase  string          `xml:"nrRecArqBase"`
	Indexistinfo  int64           `xml:"indExistInfo"`
	Infocpseg     []Infocpseg     `xml:"infoCPSeg,omitempty"`
	Infocontrib   Infocontrib     `xml:"infoContrib"`
	Ideestab      []Ideestab      `xml:"ideEstab"`
	Infocrcontrib []Infocrcontrib `xml:"infoCRContrib"`
}

type Infocpseg struct {
	Vrdesccp float64 `xml:"vrDescCP"`
	Vrcpseg  float64 `xml:"vrCpSeg"`
}

type Infocontrib struct {
	Classtrib string   `xml:"classTrib"`
	Infopj    []Infopj `xml:"infoPJ,omitempty"`
}

type Infopj struct {
	Indcoop              []string     `xml:"indCoop,omitempty"`
	Indconstr            int64        `xml:"indConstr"`
	Indsubstpatr         []string     `xml:"indSubstPatr,omitempty"`
	Percredcontrib       []string     `xml:"percRedContrib,omitempty"`
	Perctransf           []string     `xml:"percTransf,omitempty"`
	Indtribfolhapispasep []string     `xml:"indTribFolhaPisPasep,omitempty"`
	Infoatconc           []Infoatconc `xml:"infoAtConc,omitempty"`
}

type Infoatconc struct {
	Fatormes float64 `xml:"fatorMes"`
	Fator13  float64 `xml:"fator13"`
}

type Infoestabref struct {
	Aliqrat      int64    `xml:"aliqRat"`
	Fap          []string `xml:"fap,omitempty"`
	Aliqratajust []string `xml:"aliqRatAjust,omitempty"`
}

type Infocomplobra struct {
	Indsubstpatrobra int64 `xml:"indSubstPatrObra"`
}

type Infotercsusp struct {
	Codterc string `xml:"codTerc"`
}

type Basesremun struct {
	Indincid  int64       `xml:"indIncid"`
	Codcateg  int64       `xml:"codCateg"`
	Basescp   Basescp     `xml:"basesCp"`
	Basescp13 []Basescp13 `xml:"basesCp13,omitempty"`
}

type Basescp struct {
	Vrbccp00       float64  `xml:"vrBcCp00"`
	Vrbccp15       float64  `xml:"vrBcCp15"`
	Vrbccp20       float64  `xml:"vrBcCp20"`
	Vrbccp25       float64  `xml:"vrBcCp25"`
	Vrsuspbccp00   float64  `xml:"vrSuspBcCp00"`
	Vrsuspbccp15   float64  `xml:"vrSuspBcCp15"`
	Vrsuspbccp20   float64  `xml:"vrSuspBcCp20"`
	Vrsuspbccp25   float64  `xml:"vrSuspBcCp25"`
	Vrbccp00va     []string `xml:"vrBcCp00VA,omitempty"`
	Vrbccp15va     []string `xml:"vrBcCp15VA,omitempty"`
	Vrbccp20va     []string `xml:"vrBcCp20VA,omitempty"`
	Vrbccp25va     []string `xml:"vrBcCp25VA,omitempty"`
	Vrsuspbccp00va []string `xml:"vrSuspBcCp00VA,omitempty"`
	Vrsuspbccp15va []string `xml:"vrSuspBcCp15VA,omitempty"`
	Vrsuspbccp20va []string `xml:"vrSuspBcCp20VA,omitempty"`
	Vrsuspbccp25va []string `xml:"vrSuspBcCp25VA,omitempty"`
	Vrdescsest     float64  `xml:"vrDescSest"`
	Vrcalcsest     float64  `xml:"vrCalcSest"`
	Vrdescsenat    float64  `xml:"vrDescSenat"`
	Vrcalcsenat    float64  `xml:"vrCalcSenat"`
	Vrsalfam       float64  `xml:"vrSalFam"`
	Vrsalmat       float64  `xml:"vrSalMat"`
}

type Basescp13 struct {
	Vrbccp00     float64 `xml:"vrBcCp00"`
	Vrbccp15     float64 `xml:"vrBcCp15"`
	Vrbccp20     float64 `xml:"vrBcCp20"`
	Vrbccp25     float64 `xml:"vrBcCp25"`
	Vrsuspbccp00 float64 `xml:"vrSuspBcCp00"`
	Vrsuspbccp15 float64 `xml:"vrSuspBcCp15"`
	Vrsuspbccp20 float64 `xml:"vrSuspBcCp20"`
	Vrsuspbccp25 float64 `xml:"vrSuspBcCp25"`
}

type Basesavnport struct {
	Vrbccp00 float64 `xml:"vrBcCp00"`
	Vrbccp15 float64 `xml:"vrBcCp15"`
	Vrbccp20 float64 `xml:"vrBcCp20"`
	Vrbccp25 float64 `xml:"vrBcCp25"`
	Vrbccp13 float64 `xml:"vrBcCp13"`
	Vrdesccp float64 `xml:"vrDescCP"`
}

type Basesaquis struct {
	Indaquis    int64   `xml:"indAquis"`
	Vlraquis    float64 `xml:"vlrAquis"`
	Vrcpdescpr  float64 `xml:"vrCPDescPR"`
	Vrcpnret    float64 `xml:"vrCPNRet"`
	Vrratnret   float64 `xml:"vrRatNRet"`
	Vrsenarnret float64 `xml:"vrSenarNRet"`
	Vrcpcalcpr  float64 `xml:"vrCPCalcPR"`
	Vrratdescpr float64 `xml:"vrRatDescPR"`
	Vrratcalcpr float64 `xml:"vrRatCalcPR"`
	Vrsenardesc float64 `xml:"vrSenarDesc"`
	Vrsenarcalc float64 `xml:"vrSenarCalc"`
}

type Basescomerc struct {
	Indcomerc   int64    `xml:"indComerc"`
	Vrbccompr   float64  `xml:"vrBcComPR"`
	Vrcpsusp    []string `xml:"vrCPSusp,omitempty"`
	Vrratsusp   []string `xml:"vrRatSusp,omitempty"`
	Vrsenarsusp []string `xml:"vrSenarSusp,omitempty"`
}

type Infocrestab struct {
	Tpcr     int64    `xml:"tpCR"`
	Vrcr     float64  `xml:"vrCR"`
	Vrsuspcr []string `xml:"vrSuspCR,omitempty"`
}

type Basespispasep struct {
	Vrbcpispasep     float64 `xml:"vrBcPisPasep"`
	Vrbcpispasepsusp float64 `xml:"vrBcPisPasepSusp"`
}

type S5012 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evtirrf struct {
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Infoirrf      Infoirrf      `xml:"infoIRRF"`
}

type Infoirrf struct {
	Nrrecarqbase string      `xml:"nrRecArqBase"`
	Indexistinfo int64       `xml:"indExistInfo"`
	Infocrmen    []Infocrmen `xml:"infoCRMen"`
	Infocrdia    []Infocrdia `xml:"infoCRDia"`
}

type Infocrmen struct {
	Crmen   string  `xml:"CRMen"`
	Vrcrmen float64 `xml:"vrCRMen"`
}

type Infocrdia struct {
	Perapurdia int64   `xml:"perApurDia"`
	Crdia      string  `xml:"CRDia"`
	Vrcrdia    float64 `xml:"vrCRDia"`
}

type S5013 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evtfgts struct {
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Infofgts      Infofgts      `xml:"infoFGTS"`
}

type S5501 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evttribproctrab struct {
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Ideproc       Ideproc       `xml:"ideProc"`
}

type Infotributos struct {
	Perref        string          `xml:"perRef"`
	Infocrcontrib []Infocrcontrib `xml:"infoCRContrib"`
}

type S5503 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evtfgtsproctrab struct {
	Id             string         `xml:"Id"`
	Ideevento      Ideevento      `xml:"ideEvento"`
	Ideempregador  Ideempregador  `xml:"ideEmpregador"`
	Ideproc        Ideproc        `xml:"ideProc"`
	Idetrabalhador Idetrabalhador `xml:"ideTrabalhador"`
	Infotrabfgts   []Infotrabfgts `xml:"infoTrabFGTS"`
}

type Infofgtsproctrab struct {
	Totalfgts float64    `xml:"totalFGTS"`
	Ideestab  []Ideestab `xml:"ideEstab,omitempty"`
}

type Baseperref struct {
	Perref          string   `xml:"perRef"`
	Codcateg        int64    `xml:"codCateg"`
	Tpvalorproctrab int64    `xml:"tpValorProcTrab"`
	Remfgtsproctrab float64  `xml:"remFGTSProcTrab"`
	Dpsfgtsproctrab []string `xml:"dpsFGTSProcTrab,omitempty"`
	Remfgtssefip    []string `xml:"remFGTSSefip,omitempty"`
	Dpsfgtssefip    []string `xml:"dpsFGTSSefip,omitempty"`
	Remfgtsdecant   []string `xml:"remFGTSDecAnt,omitempty"`
	Dpsfgtsdecant   []string `xml:"dpsFGTSDecAnt,omitempty"`
}

type S8200 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evtanotjud struct {
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Infoprocesso  Infoprocesso  `xml:"infoProcesso"`
	Infoanotjud   Infoanotjud   `xml:"infoAnotJud"`
}

type Infoanotjud struct {
	Cpftrab       string         `xml:"cpfTrab"`
	Nmtrab        string         `xml:"nmTrab"`
	Dtnascto      string         `xml:"dtNascto"`
	Dtadm         string         `xml:"dtAdm"`
	Matricula     string         `xml:"matricula"`
	Codcateg      int64          `xml:"codCateg"`
	Natatividade  int64          `xml:"natAtividade"`
	Tpcontr       int64          `xml:"tpContr"`
	Dtterm        []string       `xml:"dtTerm,omitempty"`
	Tpinsctrab    []string       `xml:"tpInscTrab,omitempty"`
	Localtrabalho []string       `xml:"localTrabalho,omitempty"`
	Tpregtrab     int64          `xml:"tpRegTrab"`
	Tpregprev     int64          `xml:"tpRegPrev"`
	Cargo         []Cargo        `xml:"cargo"`
	Remuneracao   []Remuneracao  `xml:"remuneracao"`
	Incorporacao  []Incorporacao `xml:"incorporacao"`
	Afastamento   []Afastamento  `xml:"afastamento,omitempty"`
	Desligamento  []Desligamento `xml:"desligamento,omitempty"`
}

type Cargo struct {
	Dtcargo  string `xml:"dtCargo"`
	Cbocargo string `xml:"CBOCargo"`
}

type Incorporacao struct {
	Tpinsc    []string `xml:"tpInsc,omitempty"`
	Nrinsc    []string `xml:"nrInsc,omitempty"`
	Matincorp string   `xml:"matIncorp"`
}

type S8299 struct {
	XMLName xml.Name `xml:"eSocial"`
	Esocial Esocial  `xml:"eSocial"`
}

type Evtbaixa struct {
	Id            string        `xml:"Id"`
	Ideevento     Ideevento     `xml:"ideEvento"`
	Ideempregador Ideempregador `xml:"ideEmpregador"`
	Idevinculo    Idevinculo    `xml:"ideVinculo"`
	Infobaixa     Infobaixa     `xml:"infoBaixa"`
}

type Infobaixa struct {
	Mtvdeslig    string   `xml:"mtvDeslig"`
	Dtdeslig     string   `xml:"dtDeslig"`
	Dtprojfimapi []string `xml:"dtProjFimAPI,omitempty"`
	Nrproctrab   string   `xml:"nrProcTrab"`
	Observacao   []string `xml:"observacao,omitempty"`
}
