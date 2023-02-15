package inifile

import (
	"errors"
	"github.com/lantaris/rest-go-sdk/logger"
	"gopkg.in/ini.v1"
	"os"
)

type TIniFileData struct {
	Param    string
	Default  string
	Value    *ini.Key
	Required bool
}

type TIniFileSection struct {
	Section string
	Params  []TIniFileData
}

var (
	iniFileData *[]TIniFileSection
	iniFile     *ini.File
	iniFileName string
)

// ***********************************************************************
func arrayHasKey(Data []string, Key string) bool {
	for _, DataItem := range Data {
		if DataItem == Key {
			return true
		}
	}
	return false
}

// ***********************************************************************
func createFile(FileName string) error {
	var (
		err  error = nil
		file *os.File
	)
	file, err = os.Create(FileName)
	if err != nil {
		logger.Errorln(err)
		return err
	}
	defer file.Close()
	return nil
}

// ***********************************************************************
func fillIniFileData() error {
	var (
		err error = nil
	)
	// Iterating sections
	for SecInd := 0; SecInd < len(*iniFileData); SecInd++ {
		// Iterating parameters
		for ParamInd := 0; ParamInd < len((*iniFileData)[SecInd].Params); ParamInd++ {
			if !iniFile.Section((*iniFileData)[SecInd].Section).HasKey((*iniFileData)[SecInd].Params[ParamInd].Param) {
				logger.Traceln("Create key (" + (*iniFileData)[SecInd].Section + ":" +
					(*iniFileData)[SecInd].Params[ParamInd].Param + ")")
				_, err = iniFile.Section((*iniFileData)[SecInd].Section).NewKey((*iniFileData)[SecInd].Params[ParamInd].Param,
					(*iniFileData)[SecInd].Params[ParamInd].Default)
				if err != nil {
					logger.Errorln("Error creating section key: " + err.Error())
					return err
				}
			}
		}
	}

	err = iniFile.SaveTo(iniFileName)
	return err
}

// ***********************************************************************
func getDataParams(Section string) []string {
	var params []string = nil
	for SecInd := 0; SecInd < len(*iniFileData); SecInd++ {
		// Iterating parameters
		for ParamInd := 0; ParamInd < len((*iniFileData)[SecInd].Params); ParamInd++ {
			params = append(params, (*iniFileData)[SecInd].Params[ParamInd].Param)
		}
	}
	return params
}

// ***********************************************************************
func checkDeprications() {
	var (
		DataParams []string = nil
	)
	for _, IniSect := range iniFile.SectionStrings() {
		DataParams = getDataParams(IniSect)
		for _, IniParam := range iniFile.Section(IniSect).KeyStrings() {
			if !arrayHasKey(DataParams, IniParam) {
				logger.Debugln("Ini file parameter: [" + IniSect + ":" + IniParam +
					"] syntax error or deprecated.")
			}
		}
	}
}

// ***********************************************************************
func checkIniFileData() error {
	var (
		err error = nil
	)
	// Iterating sections
	for SecInd := 0; SecInd < len(*iniFileData); SecInd++ {
		// Iterating parameters
		for ParamInd := 0; ParamInd < len((*iniFileData)[SecInd].Params); ParamInd++ {
			if !iniFile.Section((*iniFileData)[SecInd].Section).HasKey((*iniFileData)[SecInd].Params[ParamInd].Param) {
				if (*iniFileData)[SecInd].Params[ParamInd].Required {
					if (*iniFileData)[SecInd].Params[ParamInd].Default == "" {
						logger.Errorln("Parameter (" + (*iniFileData)[SecInd].Section + ":" +
							(*iniFileData)[SecInd].Params[ParamInd].Param + ") not set")
						return errors.New("INI param not found")
					}
				}

				logger.Traceln("Create key (" + (*iniFileData)[SecInd].Params[ParamInd].Param + ")")
				_, err = iniFile.Section((*iniFileData)[SecInd].Section).NewKey((*iniFileData)[SecInd].Params[ParamInd].Param,
					(*iniFileData)[SecInd].Params[ParamInd].Default)
				if err != nil {
					logger.Errorln("Error creating section key: " + err.Error())
					return err
				}
			}
			(*iniFileData)[SecInd].Params[ParamInd].Value = iniFile.Section((*iniFileData)[SecInd].Section).Key((*iniFileData)[SecInd].Params[ParamInd].Param)
		}
	}

	err = Save()
	if err != nil {
		logger.Errorln("Error save ini file" + err.Error())
		return err
	}
	return err
}

// ***********************************************************************
func Save() error {
	var err error = nil
	err = iniFile.SaveTo(iniFileName)
	if err != nil {
		logger.Errorln("Error saving data to ini file: " + err.Error())
		return err
	}
	return err
}

// ***********************************************************************
func Initialization(FileName string, IniData *[]TIniFileSection) error {
	var (
		err         error = nil
		newfileFlag bool  = false
	)

	// Pointer for IniFileData
	iniFileData = IniData
	iniFileName = FileName

	// Check file exist
	if _, err = os.Stat(iniFileName); err != nil {
		logger.Debugln("Ini file not exist. Creating default ini file: " + iniFileName)
		err = createFile(iniFileName)
		newfileFlag = true
		if err != nil {
			logger.Errorln("Error creating ini file." + err.Error())
			return err
		}
	}

	// Open ini file
	iniFile, err = ini.Load(iniFileName)
	if err != nil {
		logger.Errorln("Error opening ini file (" + iniFileName + " ):" + err.Error())
		return err
	}

	// Fill data if new file
	if newfileFlag {
		err = fillIniFileData()
		if err != nil {
			return err
		}
	}

	// Check ini data
	err = checkIniFileData()
	if err != nil {
		return err
	}

	// Check ini data
	checkDeprications()

	return err
}

// ***********************************************************************
func GetParam(Section string, Param string) *ini.Key {
	for SecInd := 0; SecInd < len(*iniFileData); SecInd++ {
		// Iterating parameters
		if (*iniFileData)[SecInd].Section == Section {
			for ParamInd := 0; ParamInd < len((*iniFileData)[SecInd].Params); ParamInd++ {
				if (*iniFileData)[SecInd].Params[ParamInd].Param == Param {
					return (*iniFileData)[SecInd].Params[ParamInd].Value
				}
			}
		}
	}
	return nil
}

// ***********************************************************************
func SetParam(Section string, Param string, Value string) error {
	var err error = nil
	logger.Errorln("Creating/change ini parameter [" + Section + ":" + Param + "]")
	_, err = iniFile.Section(Section).NewKey(Param, Value)
	if err != nil {
		logger.Errorln("Error creating section key: " + err.Error())
		return err
	}
	//err = Save()
	return err
}
