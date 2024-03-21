package routes

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func GetConvertFiles(w http.ResponseWriter, r *http.Request) {
	// Extraer el nombre del archivo de los parámetros de la consulta
	fileName := r.URL.Query().Get("fileName")

	if fileName == "" {
		http.Error(w, "Parámetro 'fileName' faltante en la solicitud", http.StatusBadRequest)
		return
	}

	// Definir las rutas de archivo de entrada y salida

	fileNames := cargarNombresDeArchivoDesdeEnv()

	// Imprime los nombres de archivo
	fmt.Println("Nombres de archivo cargados desde el archivo .env:")
	var nombresValidos []string
	for _, fileName := range fileNames {
		if validarNombreArchivo(fileName) {
			nombresValidos = append(nombresValidos, fileName)
			fmt.Println("-------------------", nombresValidos)
			break

		}
		fmt.Println(fileName)
	}

	// Responder con un mensaje de éxito
	fmt.Fprintf(w, "Archivo transformado correctamente: %s", fileName)
}

type CSV struct {
	inputFilePath  string
	outputFilePath string
	firstLine      string
}

func NewCSV(inputFilePath, outputFilePath, firstLine string) *CSV {
	return &CSV{inputFilePath: inputFilePath, outputFilePath: outputFilePath, firstLine: firstLine}
}

func (c *CSV) TransformCSV() {
	content, err := c.readFileContent(c.inputFilePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	indexOfFirstLineBreak := strings.Index(content, "\n")
	if indexOfFirstLineBreak != -1 {
		content = content[indexOfFirstLineBreak+1:]
	}

	content = c.firstLine + "\n" + content
	modifiedContent := c.replaceCommasAndQuotes(content)

	err = ioutil.WriteFile(c.outputFilePath, []byte(modifiedContent), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Transformation completed. Result written to:", c.outputFilePath)
}

func (c *CSV) readFileContent(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (c *CSV) replaceCommasAndQuotes(input string) string {
	var result strings.Builder
	insideQuotes := false
	for _, char := range input {
		if char == '"' {
			insideQuotes = !insideQuotes
		} else if insideQuotes && char == ',' {
			result.WriteRune('.')
		} else if char != '"' {
			result.WriteRune(char)
		}
	}
	return result.String()
}

func cargarNombresDeArchivoDesdeEnv() []string {
	// Lee el contenido del archivo .env
	file, err := os.Open(".env")
	if err != nil {
		log.Fatalf("Error al abrir el archivo .env: %v", err)
	}
	defer file.Close()

	// Lee la línea del archivo que contiene los nombres de archivo
	var envLine string
	_, err = fmt.Fscanf(file, "%s", &envLine)
	if err != nil {
		log.Fatalf("Error al leer el archivo .env: %v", err)
	}

	// Obtén los nombres de archivo de la línea leída
	nombresDeArchivo := strings.Split(strings.Split(envLine, "=")[1], ",")

	return nombresDeArchivo
}
func validarNombreArchivo(fileName string) bool {
	inputFilePath := "C:/data/" + fileName
	outputFilePath := "C:/data/" + fileName

	fmt.Println("------el nombre del archivo es-------------", fileName)

	switch {
	case strings.HasSuffix(fileName, "scotiabank_wm_callback_to_vm.csv"):
		fmt.Println("Procesando scotiabank_wm_callback_to_vm.csv")
		firstLine := " Name,Surveyid, Email, Invitation_date, Responsedate, All_Segments, Number_of_callbacks_attempted, Closure_Status,    Customer_Issue_Pending, Customer_Issue_Raised, Customer_Suggestions_Follow,    Outcome, Outcome_Note, Customer_Experience_Triggers_Note,     Customer_Experience_Triggers, Select_feedback_for_Huddles_, Investment_Specialist, Age_Group, Assests_Under_Administration_Integer,   Assets_Under_Administration_Bucket, Customer_Branch_Region_Name,Callback_Role, Client_City, Client_Id, Client_Open_Date, Client_Province, Client_Street_Address, Customer_Branch_District_Id, Customer_Branch_District_Name, Customer_Branch_Id, Customer_Branch, Customer_Branch_Region_Id, Investable_Asset_Tier, Client_Postal_Code, Reference_ID, Risk_Tolerance, Share_of_Wallet_Bucket, Share_of_Wallet_Fraction, Tenure, Development_Cycle, Survey_Type, Country_Description, Country, Customer_Country, Customer_ID, First_Name, Last_Name, Business_Phone, Customer_Preferred_Language, Survey_Method, Investment_Specialist_ID, Survey_Status, Investment_Specialist_Branch, Branch_Banking_Territory, Branch_Banking_Region, Head_ID, Head, Region_ID, Region, Territory_ID, Territory, Invitation_Date_EST, Response_Date_EST, Additional_Feedback, Investment_Advisor_OSAT, Why_WM_Other_Banks, Consolidating_accounts, BanBif_LTR, BanBif, Banco_de_Chile_LTR, Banco_de_Chile, BBVA_LTR, BBVA, BCI_LTR, BCI, BCP_LTR, BCP, BICE_LTR, BICE, BMO_LTR, BMO, Expected_Business, CIBC_LTR, CIBC, Client_Value, Fulfilled_Commitments, Relationship_Drivers_Comment, Ease_of_Contact, Effectiveness, Partner_Engagement, Investment_Specialist_OSAT_Comment, None, Flexible_survey_length_chosen, Intelligo_LTR, Intelligo, Itau_LTR, Itau, Do_you_know_your_Scotiabank_Investment_Banker, Know_My_Business, LTR_Comment, Relationship_LTR, None2, OpsOpsOther_Financial_Institution_Name, Other_Financial_Institution, Other_Financial_Institutions_Comment, Other_LTR, Other_Financial_Institution2, Other, Financial_plan_review, Portfolio_Performance_Investment_Solutions, Is_Scotiabank_your_Primary_Banking_Partner, Timely_Advice, RBC_LTR, RBC, Investment_Specialists_Contact_Frequency, Retirement, Santander_LTR, Santander, Education_savings, Security_LTR, Security, Investment_Specialist_OSAT, TD_LTR, TD, Relationship_Trust, Understanding_Needs, Advice_and_Solutions, Callback_Updated_EST, Callback_by, Accival_Citibanamex, Citibanamex_(Banca_Patrimonial_Privada), BBVA_Banca_Patrimonial_Privada, Banorte_IXE_(Banca_Patrimonial_Privada), HSBC_Banca_Patrimonial_Privada, Inbursa_(Banca_Patrimonial_Privada), Santander_(Banca_Patrimonial_Privada), UBS, Bank_Of_America, JP_Morgan, Credit_Suisse, Goldman_Sachs, Accival_Citibanamex_LTR, Citibanamex_Banca_Patrimonial_Privada_LTR, BBVA_Banca_Patrimonial_Privada_LTR, Banorte_IXE_Banca_Patrimonial_Privada, HSBC_Banca_Patrimonial_Privada_LTR, Inbursa_Banca_Patrimonial_Privada, Santander_Banca_Patrimonial_Privada_LTR, UBS_LTR, Bank_Of_America_LTR, JP_Morgan_LTR, Credit_Suisse_LTR, Goldman_Sachs_LTR, _MEXICO_Secondary_Is_Scotiabank_your_Primary_Banking_Partner"
		csv := NewCSV(inputFilePath, outputFilePath, firstLine)
		csv.TransformCSV()
		return true
	case strings.HasSuffix(fileName, "scotiabank_wm_response_to_vm.csv"):
		fmt.Println("Procesando scotiabank_wm_response_to_vm.csv")
		firstLine := "Name,SurveyID,AgeGroup,AllSegments,AssetsUnderAdministrationInteger,AssetsUnderAdministrationBucket,BusinessPhone,ClientCity,ClientId,ClientStreetAddress,ClientOpenDate,ClientPostalCode,ClientProvince,Country,CountryDescription,CustomerCountry,CustomerID,Email,InvestableAssetTier,InvestmentSpecialist,InvitationDate,ReferenceID,RiskTolerance,ShareOfWalletBucket,ShareOfWalletFraction,Tenure,InvitationDateEST,CustomerBranch,CustomerBranchDistrictId,CustomerBranchDistrictName,CustomerBranchId,CustomerBranchRegionId,CustomerBranchRegionName,CustomerPreferredLanguage,DevelopmentCycle,FirstName,LastName,SurveyMethod,SurveyType,Head,HeadID,InvestmentSpecialistID,Region,RegionID,Territory,TerritoryID,ResponseDate,BranchBankingRegion,Branch,BranchBankingTerritory,InvestmentSpecialistBranch,ResponseDateEST,SurveyStatus,KnowYourScotiabankInvestmentBanker,InvestmentAdvisorOSAT,LTRComment,RelationshipLTR,AdviceAndSolutions,ClientValue,EaseOfContact,FlexibleSurveyLengthChosen,FulfilledCommitments,nvestmentSpecialistOSAT,InvestmentSpecialistOSATComment,InvestmentSpecialistsContactFrequency,IsScotiabankYourPrimaryBankingPartner,KnowMyBusiness,PartnerEngagement,PortfolioPerformanceInvestmentSolutions,UnderstandingNeeds,BanBif,BBVA,BCP,BMO,CIBC,Intelligo,OtherFinancialInstitution,RBC,TD,BanBifLTR,BBVALTR,BCPLTR,BMOLTR,CIBCLTR,ConsolidatingAccounts,EducationSavings,FinancialPlanReview,IntelligoLTR,None,Other,OtherFinancialInstitution2,OpsOpsOtherFinancialInstitutionName,OtherFinancialInstitutionsComment,OtherLTR,RBC_LTR,RelationshipDriversComment,Retirement,TD_LTR,AdditionalFeedback,BancoDeChile,BancoDeChileLTR,BCI,BCILTR,BICE,BICELTR,Effectiveness,ExpectedBusiness,Itau,ItauLTR,RelationshipTrust,Santander,SantanderLTR,Security,SecurityLTR,TimelyAdvice,Why_WM_OtherBanks,AccivalCitibanamex,CitibanamexBancaPatrimonialPrivada,BBVABancaPatrimonialPrivada,BanorteIXEBancaPatrimonialPrivada,HSBCBancaPatrimonialPrivada,InbursaBancaPatrimonialPrivada,SantanderBancaPatrimonialPrivada,UBS,BankOfAmerica,JPMorgan,CreditSuisse,GoldmanSachs,AccivalCitibanamexLTR,CitibanamexBancaPatrimonialPrivadaLTR,BBVABancaPatrimonialPrivadaLTR,BanorteIXEBancaPatrimonialPrivadaLTR,HSBCBancaPatrimonialPrivadaLTR,InbursaBancaPatrimonialPrivadaLTR,SantanderBancaPatrimonialPrivadaLTR,UBSLTR,BankOfAmericaLTR,JPMorganLTR,CreditSuisseLTR,GoldmanSachsLTR,MexicoSecondaryIsScotiabankYourPrimaryBankingPartner"
		csv := NewCSV(inputFilePath, outputFilePath, firstLine)
		csv.TransformCSV()
		return true
	case strings.HasSuffix(fileName, "scotiabank_wm_invitations_to_vm.csv"):
		fmt.Println("Procesando scotiabank_wm_invitations_to_vm.csv")
		firstLine := "Surveyid,Name,Client_Id,Client_Open_Date,Email,Investment_Specialist,Invitation_date,Reference_ID,Invitation_Date_EST,All_Segments,Assests_Under_Administration_Integer,Assets_Under_Administration_Bucket,Business_Phone,Client_City,Client_Street_Address,Client_Postal_Code,Client_Province,Country,Country_Description,Customer_Country,Investable_Asset_Tier,Risk_Tolerance,Share_of_Wallet_Bucket,Share_of_Wallet_Fraction,Age_Group,Customer_Branch,Customer_Branch_District_Id,Customer_Branch_District_Name,Customer_Branch_Id,Customer_Branch_Region_Id,Customer_Branch_Region_Name,Customer_ID,Customer_Preferred_Language,Development_Cycle,First_Name,Last_Name,Survey_Method,Survey_Type,Tenure,Branch_Banking_Region,Branch_Banking_Territory,Head,Head_ID,Investment_Specialist_Branch,Investment_Specialist_ID,Region,Region_ID,Survey_Status,Territory,Territory_ID"
		csv := NewCSV(inputFilePath, outputFilePath, firstLine)
		csv.TransformCSV()
		return true
	case strings.HasSuffix(fileName, "scotiabank_b2b_callback_to_vm.csv"):
		fmt.Println("Procesando scotiabank_b2b_callback_to_vm.csv")
		csv := NewCSV(inputFilePath, outputFilePath, firstLine)
		csv.TransformCSV()
		return true
	case strings.HasSuffix(fileName, "scotiabank_b2b_digital_inline_to_vm.csv"):
		fmt.Println("Procesando scotiabank_b2b_digital_inline_to_vm.csv")
		csv := NewCSV(inputFilePath, outputFilePath, firstLine)
		csv.TransformCSV()
		return true
	case strings.HasSuffix(fileName, "scotiabank_b2b_responses_to_vm.csv"):
		fmt.Println("Procesando scotiabank_b2b_responses_to_vm.csv")
		csv := NewCSV(inputFilePath, outputFilePath, firstLine)
		csv.TransformCSV()
		return true
	case strings.HasSuffix(fileName, "scotiabank_b2b_invitations_to_vm.csv"):
		fmt.Println("Procesando scotiabank_b2b_invitations_to_vm.csv")
		csv := NewCSV(inputFilePath, outputFilePath, firstLine)
		csv.TransformCSV()
		return true
	case strings.HasSuffix(fileName, "scotiabank_cpulse_response_to_vm.csv"):
		fmt.Println("Procesando scotiabank_cpulse_response_to_vm.csv")
		csv := NewCSV(inputFilePath, outputFilePath, firstLine)
		csv.TransformCSV()
		return true
	case strings.HasSuffix(fileName, "scotiabank_cpulse_invitations_to_vm.csv"):
		fmt.Println("Procesando scotiabank_cpulse_invitations_to_vm.csv")
		csv := NewCSV(inputFilePath, outputFilePath, firstLine)
		csv.TransformCSV()
		return true
	case strings.HasSuffix(fileName, "scotiabank_optout_to_vm.csv"):
		fmt.Println("Procesando scotiabank_optout_to_vm.csv")
		csv := NewCSV(inputFilePath, outputFilePath, firstLine)
		csv.TransformCSV()
		return true
	case strings.HasSuffix(fileName, "sb_insurance_cardif_callback_to_vm.csv"):
		fmt.Println("Procesando sb_insurance_cardif_callback_to_vm.csv")
		csv := NewCSV(inputFilePath, outputFilePath, firstLine)
		csv.TransformCSV()
		return true
	case strings.HasSuffix(fileName, "sb_insurance_cardif_Invitations_to_vm.csv"):
		fmt.Println("Procesando sb_insurance_cardif_Invitations_to_vm.csv")
		csv := NewCSV(inputFilePath, outputFilePath, firstLine)
		csv.TransformCSV()
		return true
	case strings.HasSuffix(fileName, "sb_insurance_cardif_responses_to_vm.csv"):
		fmt.Println("Procesando sb_insurance_cardif_responses_to_vm.csv")
		csv := NewCSV(inputFilePath, outputFilePath, firstLine)
		csv.TransformCSV()
		return true
	case strings.HasSuffix(fileName, ".docx"):
		return true
	case strings.HasSuffix(fileName, ".jpg"):
		return true
	default:
		fmt.Println("------no hizo nada-------------")
		return false
	}
}
