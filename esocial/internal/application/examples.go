package application

// This file contains examples of how to use the XMLBuilderService
// to generate eSocial XML events for different domains and use cases.

/*
EXAMPLE 1: Building Company Information Event (S-1000)

Create a new XMLBuilderService and use it to build company registration XML:

	service := NewXMLBuilderService()

	emp := &empresa.Empresa{
		CNPJ:                    "12345678000123",
		RazaoSocial:             "Technology Solutions Ltda",
		NomeFantasia:            stringPtr("TechSol"),
		ClassificacaoTributaria: "01",
		IndCooperativa:          "N",
		IndConstrutora:          "N",
		IndDesoneracao:          "N",
	}

	params := &EventParams{
		EmpresaCNPJ: "12345678000123",
		TpAmb:       1,
		ProcEmi:     1,
		VerProc:     "1.0.0",
		TpInsc:      1,
	}
	xmlStr, err := service.BuildEmpresaInfoXML(emp, params)
	if err != nil {
		log.Fatal(err)
	}
	// xmlStr contains the complete XML with proper formatting

EXAMPLE 2: Worker Admission Event (S-2200)

Register a new employee in the system:

	service := NewXMLBuilderService()

	worker := &trabalhador.Trabalhador{
		CPF:               "12345678901",
		ENIS:              "12345678901234567890123456789012",
		Nome:              "Maria da Silva",
		Sexo:              "F",
		DataNascimento:    time.Date(1985, 7, 15, 0, 0, 0, 0, time.UTC),
		RacaCor:           stringPtr("3"), // Parda
		GrauInstrucao:     stringPtr("08"), // Higher education
		PaisNascimento:    "105", // Brazil
		PaisNacionalidade: "105",
	}

	params := &EventParams{
		EmpresaCNPJ: "12345678000123",
		TpAmb:       1,
		ProcEmi:     1,
		VerProc:     "1.0.0",
		TpInsc:      1,
	}
	xmlStr, err := service.BuildTrabalhadorAdmissionXML(worker, params)
	if err != nil {
		log.Fatal(err)
	}

EXAMPLE 3: Worker Dismissal Event (S-2280)

Record an employee termination:

	service := NewXMLBuilderService()

	worker := &trabalhador.Trabalhador{
		CPF:  "12345678901",
		Nome: "João Santos",
	}

	dismissalDate := time.Date(2025, 3, 31, 0, 0, 0, 0, time.UTC)
	params := &WorkerDismissalParams{
		EventParams: EventParams{
			EmpresaCNPJ: "12345678000123",
			TpAmb:       1,
			ProcEmi:     1,
			VerProc:     "1.0.0",
			TpInsc:      1,
		},
		Worker:             worker,
		DtDeslig:           dismissalDate,
		MotivoDesligamento: "pedido_para_empresa",
	}
	xmlStr, err := service.BuildWorkerDismissalXML(params)
	if err != nil {
		log.Fatal(err)
	}

EXAMPLE 4: Wage Table Registration (S-1030)

Register wage types used in the company:

	service := NewXMLBuilderService()

	rubrica := &empresa.Rubrica{
		Codigo:      "001",
		Descricao:   "Salário Base",
		TipoRubrica: "1", // Normal rubric
	}

	params := &RubricaTableParams{
		EventParams: EventParams{
			EmpresaCNPJ: "12345678000123",
			TpAmb:       1,
			ProcEmi:     1,
			VerProc:     "1.0.0",
			TpInsc:      1,
		},
		Rubrica: rubrica,
	}
	xmlStr, err := service.BuildRubricaTableXML(params)
	if err != nil {
		log.Fatal(err)
	}

EXAMPLE 5: Work Location Registration (S-1040)

Register the locations where workers are allocated:

	service := NewXMLBuilderService()

	location := &ambiente_trabalho.AmbienteTrabalho{
		Codigo:          "LOC001",
		Descricao:       "Escritório Matriz",
		LocalAmbiente:   "Código IBGE do município",
		TipoInscricao:   stringPtr("1"), // CNPJ
		NumeroInscricao: stringPtr("12345678000456"),
	}

	params := &WorkLocationTableParams{
		EventParams: EventParams{
			EmpresaCNPJ: "12345678000123",
			TpAmb:       1,
			ProcEmi:     1,
			VerProc:     "1.0.0",
			TpInsc:      1,
		},
		Ambiente: location,
	}
	xmlStr, err := service.BuildWorkLocationTableXML(params)
	if err != nil {
		log.Fatal(err)
	}

EXAMPLE 6: Work Schedule Registration

Define work hour patterns:

	service := NewXMLBuilderService()

	schedule := &ambiente_trabalho.Horario{
		Codigo:                 "SCHE001",
		Descricao:            "Jornada Padrão",
		TipoJornada:          "normal",
		DuracaoJornadaSemanal: intPtr(40), // 40 hours per week
		TipoIntervalo:        stringPtr("01"), // interval type
		DuracaoIntervalo:     intPtr(60), // 60 minutes
	}

	params := &WorkScheduleTableParams{
		EventParams: EventParams{
			EmpresaCNPJ: "12345678000123",
			TpAmb:       1,
			ProcEmi:     1,
			VerProc:     "1.0.0",
			TpInsc:      1,
		},
		Horario: schedule,
	}
	xmlStr, err := service.BuildWorkScheduleTableXML(params)
	if err != nil {
		log.Fatal(err)
	}

EXAMPLE 7: Worker Remuneration Event (S-2400)

Record remuneration for a pay period:

	service := NewXMLBuilderService()

	worker := &trabalhador.Trabalhador{
		CPF:  "12345678901",
		Nome: "Pedro Costa",
	}

	params := &WorkerRemunerationParams{
		EventParams: EventParams{
			EmpresaCNPJ: "12345678000123",
			TpAmb:       1,
			ProcEmi:     1,
			VerProc:     "1.0.0",
			TpInsc:      1,
		},
		Worker:  worker,
		Periodo: "202501",
	}
	xmlStr, err := service.BuildWorkerRemunerationXML(params)
	if err != nil {
		log.Fatal(err)
	}

EXAMPLE 8: Period Closure Event (S-2300)

Declare closure of a payroll period:

	service := NewXMLBuilderService()

	params := &PeriodClosureParams{
		EventParams: EventParams{
			EmpresaCNPJ: "12345678000123",
			TpAmb:       1,
			ProcEmi:     1,
			VerProc:     "1.0.0",
			TpInsc:      1,
		},
		Periodo: "202501",
	}
	xmlStr, err := service.BuildPeriodClosureXML(params)
	if err != nil {
		log.Fatal(err)
	}

EXAMPLE 9: Temporary Leave Event (S-2250)

Record temporary absence (sick leave, statutory leave, etc):

	service := NewXMLBuilderService()

	startDate := time.Date(2025, 3, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2025, 3, 5, 0, 0, 0, 0, time.UTC)

	params := &WorkerTemporaryLeaveParams{
		EventParams: EventParams{
			EmpresaCNPJ: "12345678000123",
			TpAmb:       1,
			ProcEmi:     1,
			VerProc:     "1.0.0",
			TpInsc:      1,
		},
		TrabalhadorCPF:   "12345678901",
		DtIniAfastamento: &startDate,
		DtFimAfastamento: &endDate,
		TipoAfastamento:  "01",
	}
	xmlStr, err := service.BuildWorkerTemporaryLeaveXML(params)
	if err != nil {
		log.Fatal(err)
	}

EXAMPLE 10: Work Accident Event (S-2260)

Record a work-related accident:

	service := NewXMLBuilderService()

	accidentDate := time.Date(2025, 3, 20, 10, 30, 0, 0, time.UTC)

	params := &WorkAccidentParams{
		EventParams: EventParams{
			EmpresaCNPJ: "12345678000123",
			TpAmb:       1,
			ProcEmi:     1,
			VerProc:     "1.0.0",
			TpInsc:      1,
		},
		TrabalhadorCPF: "12345678901",
		DtAcidente:     accidentDate,
	}
	xmlStr, err := service.BuildWorkAccidentXML(params)
	if err != nil {
		log.Fatal(err)
	}

EXAMPLE 11: FGTS Bases Event (S-2500)

Record FGTS calculation bases for workers:

	service := NewXMLBuilderService()

	worker := &trabalhador.Trabalhador{
		CPF:  "12345678901",
		Nome: "Ana Silva",
	}

	params := &WorkerFGTSParams{
		EventParams: EventParams{
			EmpresaCNPJ: "12345678000123",
			TpAmb:       1,
			ProcEmi:     1,
			VerProc:     "1.0.0",
			TpInsc:      1,
		},
		Worker:  worker,
		Periodo: "202501",
	}
	xmlStr, err := service.BuildWorkerFGTSXML(params)
	if err != nil {
		log.Fatal(err)
	}

HELPER FUNCTIONS

The examples above use helper functions for pointer conversion:

	func stringPtr(s string) *string {
		return &s
	}

	func intPtr(i int) *int {
		return &i
	}

Alternatively, use Go 1.17+ approach:

	idade := 40
	duração := &idade // Direct reference

MESSAGE PROCESSING FLOW

Typical flow for processing multiple events:

	service := NewXMLBuilderService()

	// Collect all events
	var events []string

	// Create common event parameters
	eventParams := EventParams{
		EmpresaCNPJ: cnpj,
		TpAmb:       1,
		ProcEmi:     1,
		VerProc:     "1.0.0",
		TpInsc:      1,
	}

	// 1. Register company
	empXML, _ := service.BuildEmpresaInfoXML(empresa, &eventParams)
	events = append(events, empXML)

	// 2. Register locations
	for _, loc := range locations {
		locParams := &WorkLocationTableParams{
			EventParams: eventParams,
			Ambiente:    loc,
		}
		locXML, _ := service.BuildWorkLocationTableXML(locParams)
		events = append(events, locXML)
	}

	// 3. Register wages
	for _, rubrica := range rubricas {
		rubParams := &RubricaTableParams{
			EventParams: eventParams,
			Rubrica:     rubrica,
		}
		rubXML, _ := service.BuildRubricaTableXML(rubParams)
		events = append(events, rubXML)
	}

	// 4. Register workers
	for _, worker := range workers {
		workerParams := &EventParams{
			EmpresaCNPJ: eventParams.EmpresaCNPJ,
			TpAmb:       eventParams.TpAmb,
			ProcEmi:     eventParams.ProcEmi,
			VerProc:     eventParams.VerProc,
			TpInsc:      eventParams.TpInsc,
		}
		workerXML, _ := service.BuildTrabalhadorAdmissionXML(worker, workerParams)
		events = append(events, workerXML)
	}

	// Send events to eSocial system
	for _, event := range events {
		// Send to eSocial Web Service
		response, err := sendToeSocial(event)
		// Handle response
	}

ERROR HANDLING BEST PRACTICES

Always check errors and implement proper error handling:

	params := &EventParams{
		EmpresaCNPJ: "12345678000123",
		TpAmb:       1,
		ProcEmi:     1,
		VerProc:     "1.0.0",
		TpInsc:      1,
	}
	xml, err := service.BuildEmpresaInfoXML(emp, params)
	if err != nil {
		switch err.(type) {
		case *xml.SyntaxError:
			log.Printf("XML syntax error: %v", err)
		case *fmt.Errorf:
			log.Printf("Formatting error: %v", err)
		default:
			log.Printf("Unknown error: %v", err)
		}
		return
	}

	// Validate XML before sending
	if !validateXML(xml) {
		log.Fatal("Generated XML is invalid")
	}

PERFORMANCE CONSIDERATIONS

- XMLBuilderService is stateless and thread-safe
- Can reuse single instance across goroutines
- Consider pooling for high-throughput scenarios
- Cache common enterprise data to reduce rebuilds

INTEGRATION WITH DATABASE

Typical integration pattern:

	func importWorkerFromDatabase(db *sql.DB, workerID string) (string, error) {
		worker := &trabalhador.Trabalhador{}

		err := db.QueryRow(
			"SELECT cpf, nome, sexo, data_nascimento, pais_nascimento, pais_nacionalidade FROM trabalhadores WHERE id = $1",
			workerID,
		).Scan(
			&worker.CPF,
			&worker.Nome,
			&worker.Sexo,
			&worker.DataNascimento,
			&worker.PaisNascimento,
			&worker.PaisNacionalidade,
		)

		if err != nil {
			return "", err
		}

		service := NewXMLBuilderService()
		return service.BuildTrabalhadorAdmissionXML(worker, "12345678000123", int64(1), int64(1), "1.0.0", int64(1))
	}
*/
