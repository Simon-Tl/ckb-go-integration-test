package mock

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/nervosnetwork/ckb-sdk-go/v2/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
)

const BASE_MOCK_URL = "http://127.0.0.1:5000/test/"

type Request struct {
	ID      string        `json:"id"`
	JSONRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type RequestParamsNil struct {
	ID      string `json:"id"`
	JSONRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
}

type Response struct {
	ID      string      `json:"id"`
	JSONRPC string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	Error   interface{} `json:"error"`
}

type MockData struct {
	Request  Request  `json:"request"`
	Response Response `json:"response"`
}

func fetchMockData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err

	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

// Parse Version from Params
func parseVersion(versionInterface interface{}) (uint32, error) {
	version, err := interfaceToUint(versionInterface)
	if err != nil {
		return 0, fmt.Errorf("parseVersion failed: %w", err)
	}
	return uint32(version), nil
}

// 公共函数：解析 cell_deps
func parseCellDeps(cellDepsInterface interface{}) ([]*types.CellDep, error) {
	celldepList := make([]*types.CellDep, 0)
	cellDepsSlice, ok := cellDepsInterface.([]interface{})
	if !ok {
		return nil, fmt.Errorf("parseCellDeps failed: expected []interface{}, got %T", cellDepsInterface)
	}
	for _, dep := range cellDepsSlice {
		if dep == nil {
			continue
		}
		celldepmap, err := interfaceToMapString(dep)
		if err != nil {
			return nil, fmt.Errorf("parseCellDeps failed: %w", err)
		}
		outpointmap, err := interfaceToMapString(celldepmap["out_point"])
		if err != nil {
			return nil, fmt.Errorf("parseCellDeps failed: %w", err)
		}
		indexer, err := interfaceToUint(outpointmap["index"])
		if err != nil {
			return nil, fmt.Errorf("parseCellDeps failed: %w", err)
		}
		outpoint := types.OutPoint{
			TxHash: types.HexToHash(outpointmap["tx_hash"].(string)),
			Index:  uint32(indexer),
		}
		celldep := types.CellDep{
			OutPoint: &outpoint,
			DepType:  types.DepType(celldepmap["dep_type"].(string)), // 假设为 dep_group 类型
		}
		celldepList = append(celldepList, &celldep)
	}
	return celldepList, nil
}

// 公共函数：解析 header_deps
func parseHeaderDeps(headerDepsInterface interface{}) ([]types.Hash, error) {
	headerDeps, err := interfaceToHashSlice(headerDepsInterface)
	if err != nil {
		return nil, fmt.Errorf("parseHeaderDeps failed: %w", err)
	}
	return headerDeps, nil
}

//公共函数: 解析 cellbase

// 公共函数：解析 Cellbase
func parseCellbase(cellbaseInterface interface{}) (*types.CellbaseTemplate, error) {
	cellbaseMap, err := interfaceToMapString(cellbaseInterface)
	if err != nil {
		return nil, fmt.Errorf("parseCellbase failed: %w", err)
	}

	hash, err := parseHash(cellbaseMap["hash"])
	if err != nil {
		return nil, fmt.Errorf("parseCellbase failed: %w", err)
	}

	cycles, err := parseCycles(cellbaseMap["cycles"])
	if err != nil {
		return nil, fmt.Errorf("parseCellbase failed: %w", err)
	}

	data, err := parseTransaction(cellbaseMap["data"])
	if err != nil {
		return nil, fmt.Errorf("parseCellbase failed: %w", err)
	}

	return &types.CellbaseTemplate{
		Hash:   hash,
		Cycles: cycles,
		Data:   data,
	}, nil
}

// 辅助函数：解析 hash 字段
func parseHash(data interface{}) (types.Hash, error) {
	hashStr, ok := data.(string)
	if !ok {
		return types.Hash{}, fmt.Errorf("parseHash failed: expected string, got %T", data)
	}
	return types.HexToHash(hashStr), nil
}

// 辅助函数：解析 cycles 字段
func parseCycles(data interface{}) (*uint64, error) {
	if data == nil {
		return nil, nil
	}
	cycles, err := interfaceToUint(data)
	if err != nil {
		return nil, fmt.Errorf("parseCycles failed: %w", err)
	}
	cyclesPtr := uint64(cycles)
	return &cyclesPtr, nil
}

func parseUncles(data interface{}) ([]types.UncleTemplate, error) {
	// 填充代码解析 Uncles 列表
	return nil, nil // 示例返回
}

func parseProposals(data interface{}) ([]string, error) {
	stringSlice, ok := data.([]interface{})
	if !ok {
		return nil, fmt.Errorf("parseProposals failed: expected []interface{}, got %T", data)
	}
	proposals := make([]string, len(stringSlice))
	for i, v := range stringSlice {
		str, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("parseProposals failed: expected string, got %T", v)
		}
		proposals[i] = str
	}
	return proposals, nil
}

func parseExtension(data interface{}) *json.RawMessage {

	if data == nil {
		return nil
	}
	raw, ok := data.(string)
	if !ok {
		return nil
	}
	jsonString := fmt.Sprintf("\"%s\"", raw)

	msg := json.RawMessage(jsonString)
	fmt.Println(msg)

	return &msg
}

func parseTransactions(data interface{}) ([]types.TransactionTemplate, error) {
	// 填充代码解析 Transactions 列表
	return nil, nil // 示例返回
}

func parseBlockTemplate(data []interface{}) (types.BlockTemplate, error) {

	dataMap, err := interfaceSliceToMapString(data)
	if err != nil {
		return types.BlockTemplate{}, fmt.Errorf("parseBlockTemplate dataMap failed: %w", err)
	}

	version, err := interfaceToUint(dataMap["version"])
	if err != nil {
		return types.BlockTemplate{}, fmt.Errorf("parseBlockTemplate version failed: %w", err)
	}

	compactTarget, err := interfaceToUint(dataMap["compact_target"])
	if err != nil {
		return types.BlockTemplate{}, fmt.Errorf("parseBlockTemplate compact_target failed: %w", err)
	}

	currentTime, err := interfaceToUint(dataMap["current_time"])
	if err != nil {
		return types.BlockTemplate{}, fmt.Errorf("parseBlockTemplate  current_time failed: %w", err)
	}

	number, err := interfaceToUint(dataMap["number"])
	if err != nil {
		return types.BlockTemplate{}, fmt.Errorf("parseBlockTemplate number failed: %w", err)
	}

	epoch, err := interfaceToUint(dataMap["epoch"])
	if err != nil {
		return types.BlockTemplate{}, fmt.Errorf("parseBlockTemplate epoch failed: %w", err)
	}

	parentHash := types.HexToHash(dataMap["parent_hash"].(string))

	cyclesLimit, err := interfaceToUint(dataMap["cycles_limit"])
	if err != nil {
		return types.BlockTemplate{}, fmt.Errorf("parseBlockTemplate cycles_limit failed: %w", err)
	}

	bytesLimit, err := interfaceToUint(dataMap["bytes_limit"])
	if err != nil {
		return types.BlockTemplate{}, fmt.Errorf("parseBlockTemplate bytes_limit failed: %w", err)
	}

	unclesCountLimit, err := interfaceToUint(dataMap["uncles_count_limit"])
	if err != nil {
		return types.BlockTemplate{}, fmt.Errorf("parseBlockTemplate failed: %w", err)
	}

	uncles, err := parseUncles(dataMap["uncles"])
	if err != nil {
		return types.BlockTemplate{}, fmt.Errorf("parseBlockTemplate failed: %w", err)
	}

	transactions, err := parseTransactions(dataMap["transactions"])
	if err != nil {
		return types.BlockTemplate{}, fmt.Errorf("parseBlockTemplate failed: %w", err)
	}

	proposals, err := parseProposals(dataMap["proposals"])
	if err != nil {
		return types.BlockTemplate{}, fmt.Errorf("parseBlockTemplate failed: %w", err)
	}

	cellbase, err := parseCellbase(dataMap["cellbase"])
	if err != nil {
		return types.BlockTemplate{}, fmt.Errorf("parseBlockTemplate failed: %w", err)
	}

	workId, err := interfaceToUint(dataMap["work_id"])
	if err != nil {
		return types.BlockTemplate{}, fmt.Errorf("parseBlockTemplate failed: %w", err)
	}

	dao := types.HexToHash(dataMap["dao"].(string))

	extension := parseExtension(dataMap["extension"])

	return types.BlockTemplate{
		Version:          uint32(version),
		CompactTarget:    uint32(compactTarget),
		CurrentTime:      uint64(currentTime),
		Number:           uint64(number),
		Epoch:            uint64(epoch),
		ParentHash:       parentHash,
		CyclesLimit:      uint64(cyclesLimit),
		BytesLimit:       uint64(bytesLimit),
		UnclesCountLimit: uint64(unclesCountLimit),
		Uncles:           uncles,
		Transactions:     transactions,
		Proposals:        proposals,
		Cellbase:         *cellbase,
		WorkId:           uint64(workId),
		Dao:              dao,
		Extension:        extension,
	}, nil
}

// 辅助函数：解析 data (Transaction)
func parseTransaction(data interface{}) (types.Transaction, error) {
	transactionMap, err := interfaceToMapString(data)
	if err != nil {
		return types.Transaction{}, fmt.Errorf("parseTransaction failed: %w", err)
	}

	cellDeps, err := parseCellDeps(transactionMap["cell_deps"])
	if err != nil {
		return types.Transaction{}, fmt.Errorf("parseTransaction failed: %w", err)
	}

	headerDeps, err := parseHeaderDeps(transactionMap["header_deps"])
	if err != nil {
		return types.Transaction{}, fmt.Errorf("parseTransaction failed: %w", err)
	}

	inputs, err := parseInputs(transactionMap["inputs"])
	if err != nil {
		return types.Transaction{}, fmt.Errorf("parseTransaction failed: %w", err)
	}

	outputs, err := parseOutputs(transactionMap["outputs"])
	if err != nil {
		return types.Transaction{}, fmt.Errorf("parseTransaction failed: %w", err)
	}

	outputsData, err := parseOutputsData(transactionMap["outputs_data"])
	if err != nil {
		return types.Transaction{}, fmt.Errorf("parseTransaction failed: %w", err)
	}

	version, err := parseVersion(transactionMap["version"])
	if err != nil {
		return types.Transaction{}, fmt.Errorf("parseTransaction failed: %w", err)
	}

	witnesses, err := parseWitnesses(transactionMap["witnesses"])
	if err != nil {
		return types.Transaction{}, fmt.Errorf("parseTransaction failed: %w", err)
	}

	return types.Transaction{
		CellDeps:    cellDeps,
		HeaderDeps:  headerDeps,
		Inputs:      inputs,
		Outputs:     outputs,
		OutputsData: outputsData,
		Version:     version,
		Witnesses:   witnesses,
	}, nil
}

// 公共函数：解析 inputs
func parseInputs(inputsInterface interface{}) ([]*types.CellInput, error) {
	inputList := make([]*types.CellInput, 0)
	inputSlice, ok := inputsInterface.([]interface{})
	if !ok {
		return nil, fmt.Errorf("parseInputs failed: expected []interface{}, got %T", inputsInterface)
	}
	for _, input := range inputSlice {
		if input == nil {
			continue
		}
		inputmap, err := interfaceToMapString(input)
		if err != nil {
			return nil, fmt.Errorf("parseInputs failed: %w", err)
		}
		previousoutputmap, err := interfaceToMapString(inputmap["previous_output"])
		if err != nil {
			return nil, fmt.Errorf("parseInputs failed: %w", err)
		}
		previousoutindex, err := interfaceToUint(previousoutputmap["index"])
		if err != nil {
			return nil, fmt.Errorf("parseInputs failed: %w", err)
		}
		previousout := types.OutPoint{
			TxHash: types.HexToHash(previousoutputmap["tx_hash"].(string)),
			Index:  uint32(previousoutindex),
		}
		since, err := interfaceToUint(inputmap["since"])
		if err != nil {
			return nil, fmt.Errorf("parseInputs failed: %w", err)
		}
		cellInput := types.CellInput{
			Since:          uint64(since),
			PreviousOutput: &previousout,
		}
		inputList = append(inputList, &cellInput)
	}
	return inputList, nil
}

func parseOutputsData(data interface{}) ([][]byte, error) {
	stringSlice, ok := data.([]interface{})
	if !ok {
		return nil, fmt.Errorf("parseOutputsData failed: expected []interface{}, got %T", data)
	}

	result := make([][]byte, len(stringSlice))
	for i, v := range stringSlice {
		str, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("parseOutputsData failed: expected string, got %T", v)
		}
		result[i] = []byte(str)
	}
	return result, nil
}

// 公共函数：解析 outputs
func parseOutputs(outputsInterface interface{}) ([]*types.CellOutput, error) {
	outputList := make([]*types.CellOutput, 0)
	outputSlice, ok := outputsInterface.([]interface{})
	if !ok {
		return nil, fmt.Errorf("parseOutputs failed: expected []interface{}, got %T", outputsInterface)
	}

	for _, output := range outputSlice {
		if output == nil {
			continue
		}

		// 将 output 转换为 map[string]interface{}
		outputmap, err := interfaceToMapString(output)
		if err != nil {
			return nil, fmt.Errorf("parseOutputs failed: %w", err)
		}

		// 处理 lock
		var lock *types.Script
		if lockData, ok := outputmap["lock"]; ok && lockData != nil {
			lockmap, err := interfaceToMapString(lockData)
			if err != nil {
				return nil, fmt.Errorf("parseOutputs failed: %w", err)
			}

			args, err := interfaceToBytes(lockmap["args"])
			if err != nil {
				return nil, fmt.Errorf("parseOutputs failed: %w", err)
			}

			lock = &types.Script{
				CodeHash: types.HexToHash(lockmap["code_hash"].(string)),
				HashType: types.ScriptHashType(lockmap["hash_type"].(string)),
				Args:     args,
			}
		} else {
			// 如果 lock 为空，返回错误或根据需求处理
			return nil, fmt.Errorf("parseOutputs failed: lock is nil for one of the outputs")
		}

		// 处理 type（可以为空）
		var outputType *types.Script
		if typeData, ok := outputmap["type"]; ok && typeData != nil {
			typesmap, err := interfaceToMapString(typeData)
			if err != nil {
				return nil, fmt.Errorf("parseOutputs failed: %w", err)
			}

			typesargs, err := interfaceToBytes(typesmap["args"])
			if err != nil {
				return nil, fmt.Errorf("parseOutputs failed: %w", err)
			}

			outputType = &types.Script{
				CodeHash: types.HexToHash(typesmap["code_hash"].(string)),
				HashType: types.ScriptHashType(typesmap["hash_type"].(string)),
				Args:     typesargs,
			}
		} else {
			// 如果 type 为空，保持为 nil
			outputType = nil
		}

		// 处理 capacity
		capacity, err := interfaceToUint(outputmap["capacity"])
		if err != nil {
			return nil, fmt.Errorf("parseOutputs failed: %w", err)
		}

		// 构造 CellOutput
		cellOutput := types.CellOutput{
			Capacity: uint64(capacity),
			Lock:     lock,
			Type:     outputType, // Type 允许为 nil
		}

		// 将 cellOutput 加入到输出列表中
		outputList = append(outputList, &cellOutput)
	}

	return outputList, nil
}

// 公共函数：解析 outputs_data
func parseOutputData(outputDataInterface interface{}) ([][]byte, error) {
	outputDataList := make([][]byte, 0)
	dataSlice, ok := outputDataInterface.([]interface{})
	if !ok {
		return nil, fmt.Errorf("parseOutputData failed: expected []interface{}, got %T", outputDataInterface)
	}
	for _, data := range dataSlice {
		bytes, err := interfaceToBytes(data)
		if err != nil {
			return nil, fmt.Errorf("parseOutputData failed: %w", err)
		}
		outputDataList = append(outputDataList, bytes)
	}
	return outputDataList, nil
}

// 公共函数：解析 witnesses
func parseWitnesses(witnessesInterface interface{}) ([][]byte, error) {
	witnessList := make([][]byte, 0)
	witnessSlice, ok := witnessesInterface.([]interface{})
	if !ok {
		return nil, fmt.Errorf("parseWitnesses failed: expected []interface{}, got %T", witnessesInterface)
	}
	for _, witness := range witnessSlice {
		bytes, err := interfaceToBytes(witness)
		if err != nil {
			return nil, fmt.Errorf("parseWitnesses failed: %w", err)
		}
		witnessList = append(witnessList, bytes)
	}
	return witnessList, nil
}

func hashSliceToInterfaceSlice(hashes []*types.Hash) []interface{} {
	var result []interface{}
	for _, hash := range hashes {
		if hash != nil {
			result = append(result, hash.String()) // 使用 types.Hash 的 String() 方法获取十六进制表示
		}
	}
	return result
}

// 反射机制获取传入结构体中的 JSON 标签
func getStructJSONKeys(structType interface{}) ([]string, error) {
	var keys []string

	// 检查传入的是否为结构体
	val := reflect.ValueOf(structType)
	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected a struct, got %T", structType)
	}

	// 获取结构体类型的反射信息
	structTypeReflect := reflect.TypeOf(structType)
	for i := 0; i < structTypeReflect.NumField(); i++ {
		field := structTypeReflect.Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag != "" {
			keys = append(keys, jsonTag)
		}
	}
	return keys, nil
}

func compareKeys(localResult map[string]interface{}, structKeys []string) error {
	var missingKeys []string
	var extraKeys []string
	keySet := make(map[string]struct{})

	// 将 structKeys 中的 key 放入 keySet
	for _, key := range structKeys {
		keySet[key] = struct{}{}
	}

	// 检查 localResult 中的 key 是否都存在于 structKeys 中
	for k := range localResult {
		if _, found := keySet[k]; !found {
			extraKeys = append(extraKeys, k) // localResult 中的 key 不在 structKeys 中
		}
	}

	// 比较 structKeys 中的 key 是否都存在于 localResult 中
	for _, key := range structKeys {
		if _, found := localResult[key]; !found {
			missingKeys = append(missingKeys, key) // structKeys 中的 key 不在 localResult 中
		}
	}

	// 如果有缺少或多余的 key，返回 error，包含缺失或多余的 key 信息
	if len(missingKeys) > 0 || len(extraKeys) > 0 {
		errMsg := "Key mismatch:"
		if len(missingKeys) > 0 {
			errMsg += fmt.Sprintf(" Missing keys: %v.", missingKeys)
		}
		if len(extraKeys) > 0 {
			errMsg += fmt.Sprintf(" Extra keys: %v.", extraKeys)
		}
		return errors.New(errMsg)
	}

	// 如果没有错误，则键完全匹配
	return nil
}

func interfaceSliceToStringSlice(data []interface{}) ([]string, error) {
	// 创建一个和data等长的string切片
	strSlice := make([]string, len(data))
	for i, v := range data {
		// 尝试将interface{}转换为string
		if str, ok := v.(string); ok {
			strSlice[i] = str
		} else {
			// 如果某个元素不是string类型，则返回错误
			return nil, fmt.Errorf("item %d is not of type string, got %T", i, v)
		}
	}
	return strSlice, nil
}

func interfaceSliceToMapString(i []interface{}) (map[string]interface{}, error) {
	// 检查切片是否为空，避免越界访问
	if len(i) == 0 {
		return nil, fmt.Errorf("input slice is empty")
	}

	var result map[string]interface{}

	// 遍历切片，找到第一个 map[string]interface{}
	for _, item := range i {
		// 类型断言，将 interface{} 转为 map[string]interface{}
		m, ok := item.(map[string]interface{})
		if ok {
			result = m
			break
		}
	}

	// 如果没有找到任何 map[string]interface{}，返回错误
	if result == nil {
		return nil, fmt.Errorf("no valid map[string]interface{} found in slice")
	}

	//// 序列化结果为 JSON 格式
	//jsonData, err := json.Marshal(result)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to serialize map to JSON: %v", err)
	//}
	//
	//fmt.Println("Serialized JSON:", string(jsonData))

	// 返回剔除非 map[string]interface{} 之后的结果
	return result, nil
}

func interfaceToMapString(i interface{}) (map[string]interface{}, error) {
	// 类型断言，将 interface{} 转为 map[string]interface{}
	ms, ok := i.([]interface{})
	if ok {
		a, _ := interfaceSliceToMapString(ms)
		return a, nil
	}
	m, ok := i.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("failed to assert type to map[string]interface{} for value: %v", i)
	}

	// 如果断言成功，返回 map
	return m, nil
}

func interfaceToHashSlice(data interface{}) ([]types.Hash, error) {
	if data == nil {
		return []types.Hash{}, nil // 返回空的 []Hash
	}

	// 尝试将 interface{} 转换为 []string
	if stringSlice, ok := data.([]string); ok {
		hashes := make([]types.Hash, len(stringSlice))
		for i, str := range stringSlice {
			if len(str) < 2 || str[:2] != "0x" {
				return nil, fmt.Errorf("item %d is not a valid hex string with '0x' prefix", i)
			}
			// 移除 "0x" 前缀并将 hex 字符串转为 []byte
			byteArray, err := hex.DecodeString(str[2:])
			if err != nil {
				return nil, fmt.Errorf("item %d is not a valid hex string: %v", i, err)
			}
			if len(byteArray) != 32 {
				return nil, fmt.Errorf("item %d does not have length 32, length is %d", i, len(byteArray))
			}
			var hash types.Hash
			copy(hash[:], byteArray)
			hashes[i] = hash
		}
		return hashes, nil
	}

	// 尝试将 interface{} 转换为 [][]byte
	if byteSlices, ok := data.([][]byte); ok {
		hashes := make([]types.Hash, len(byteSlices))
		for i, byteArray := range byteSlices {
			if len(byteArray) != 32 {
				return nil, fmt.Errorf("item %d does not have length 32, length is %d", i, len(byteArray))
			}
			var hash types.Hash
			copy(hash[:], byteArray)
			hashes[i] = hash
		}
		return hashes, nil
	}

	// 尝试将 interface{} 转换为 []interface{}
	if slice, ok := data.([]interface{}); ok {
		byteSlices := make([][]byte, len(slice))
		for i, item := range slice {
			if byteArray, ok := item.([]byte); ok {
				if len(byteArray) != 32 {
					return nil, fmt.Errorf("item %d does not have length 32, length is %d", i, len(byteArray))
				}
				byteSlices[i] = byteArray
			} else if str, ok := item.(string); ok {
				// 将字符串转换为 []byte
				if len(str) < 2 || str[:2] != "0x" {
					return nil, fmt.Errorf("item %d is not a valid hex string with '0x' prefix", i)
				}
				byteArray, err := hex.DecodeString(str[2:])
				if err != nil {
					return nil, fmt.Errorf("item %d is not a valid hex string: %v", i, err)
				}
				if len(byteArray) != 32 {
					return nil, fmt.Errorf("item %d does not have length 32, length is %d", i, len(byteArray))
				}
				byteSlices[i] = byteArray
			} else {
				return nil, fmt.Errorf("item %d is not of type []byte or string", i)
			}
		}
		hashes := make([]types.Hash, len(byteSlices))
		for i, byteArray := range byteSlices {
			var hash types.Hash
			copy(hash[:], byteArray)
			hashes[i] = hash
		}
		return hashes, nil
	}

	// 提供更详细的错误信息
	return nil, fmt.Errorf("interface{} is not of type []string, [][]byte, or []interface{}, got %T", data)
}

func interfaceToUint(data interface{}) (uint, error) {
	switch v := data.(type) {
	case uint64:
		// 如果输入已经是 uint64，确保它在 uint 范围内
		if v > uint64(^uint(0)) { // 检查是否超过 uint 的范围
			return 0, fmt.Errorf("value exceeds uint range")
		}
		return uint(v), nil
	case string:
		// 如果是字符串，判断是否以 "0x" 开头
		if strings.HasPrefix(v, "0x") {
			// 去掉 "0x" 前缀并转换为 uint64
			parsedValue, err := strconv.ParseUint(v[2:], 16, 64)
			if err != nil {
				return 0, fmt.Errorf("failed to parse hex string to uint: %v", err)
			}
			// 检查是否超过 uint 的范围
			if parsedValue > uint64(^uint(0)) {
				return 0, fmt.Errorf("hex string value exceeds uint range")
			}
			return uint(parsedValue), nil
		} else {
			// 非十六进制字符串的处理
			parsedValue, err := strconv.ParseUint(v, 10, 64)
			if err != nil {
				return 0, fmt.Errorf("failed to parse string to uint: %v", err)
			}
			// 检查是否超过 uint 的范围
			if parsedValue > uint64(^uint(0)) {
				return 0, fmt.Errorf("string value exceeds uint range")
			}
			return uint(parsedValue), nil
		}
	default:
		return 0, fmt.Errorf("unsupported type for conversion to uint: %T", data)
	}
}

func interfaceToUintSlice(data interface{}) ([]uint, error) {
	// 检查是否为空
	if data == nil {
		return []uint{}, nil // 返回空的 []uint
	}

	// 尝试将 interface{} 转换为 []interface{}
	if slice, ok := data.([]interface{}); ok {
		// 将 []interface{} 转换为 []uint
		var result []uint
		for _, item := range slice {
			switch v := item.(type) {
			case uint:
				result = append(result, v)
			case string:
				// 处理字符串形式的十六进制数
				if strings.HasPrefix(v, "0x") {
					parsedValue, err := strconv.ParseUint(v[2:], 16, 64)
					if err != nil {
						return nil, fmt.Errorf("failed to parse hex string %v: %v", v, err)
					}
					result = append(result, uint(parsedValue))
				} else {
					return nil, fmt.Errorf("item %v is not of type uint and not a valid hex string", item)
				}
			default:
				return nil, fmt.Errorf("item %v is not of type uint or a valid hex string", item)
			}
		}
		return result, nil
	}
	return nil, fmt.Errorf("interface{} is not of type []interface{}")
}

func interfaceToBytes(data interface{}) ([]byte, error) {
	switch v := data.(type) {
	case string:
		// 如果输入为 "0x"，直接返回空字节数组
		if v == "0x" {
			return []byte{}, nil
		}
		// 去掉可能的 "0x" 前缀
		if strings.HasPrefix(v, "0x") {
			v = v[2:]
		}
		// 检查剩余的字符串是否为偶数长度（每两个字符表示一个字节）
		if len(v)%2 != 0 {
			return nil, errors.New("invalid hex string: length must be even")
		}
		// 转换为字节数组
		bytes, err := hex.DecodeString(v)
		if err != nil {
			return nil, fmt.Errorf("failed to decode hex string: %v", err)
		}
		return bytes, nil
	default:
		return nil, fmt.Errorf("unsupported type: %T", data)
	}
}

func getMockRpcClient(url string) (rpc.Client, error) {

	client, err := rpc.DialContext(context.Background(), url)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func getMockRpcClientByName(testName string) (rpc.Client, MockData, error) {
	// testName -> url
	urlNames := strings.Split(testName, "/")[1:]
	urlName := strings.Join(urlNames, "/")
	mockUrl := BASE_MOCK_URL + urlName
	// client
	client, err := getMockRpcClient(mockUrl)
	if err != nil {
		return nil, MockData{}, err
	}
	// mock data
	mockData, err := fetchMockData(mockUrl)
	if err != nil {
		return client, MockData{}, err
	}
	var bodyJson MockData
	err = json.Unmarshal(mockData, &bodyJson)
	return client, bodyJson, nil
}
