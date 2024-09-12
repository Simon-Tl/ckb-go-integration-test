package mock

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nervosnetwork/ckb-sdk-go/v2/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/v2/types"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

const BASE_MOCK_URL = "http://127.0.0.1:5000/test/"

type Request struct {
	ID      string        `json:"id"`
	JSONRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
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
			DepType:  types.DepTypeDepGroup, // 假设为 dep_group 类型
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
		outputmap, err := interfaceToMapString(output)
		if err != nil {
			return nil, fmt.Errorf("parseOutputs failed: %w", err)
		}
		lockmap, err := interfaceToMapString(outputmap["lock"])
		if err != nil {
			return nil, fmt.Errorf("parseOutputs failed: %w", err)
		}
		args, err := interfaceToBytes(lockmap["args"])
		if err != nil {
			return nil, fmt.Errorf("parseOutputs failed: %w", err)
		}
		lock := types.Script{
			CodeHash: types.HexToHash(lockmap["code_hash"].(string)),
			HashType: types.ScriptHashType(lockmap["hash_type"].(string)),
			Args:     args,
		}
		capacity, err := interfaceToUint(outputmap["capacity"])
		if err != nil {
			return nil, fmt.Errorf("parseOutputs failed: %w", err)
		}
		cellOutput := types.CellOutput{
			Capacity: uint64(capacity),
			Lock:     &lock,
			Type:     nil, // 假设 Type 为 nil
		}
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
		return []types.Hash{}, nil // 返回空的 []types.Hash
	}

	// 尝试将 interface{} 转换为 [][]byte
	if byteSlices, ok := data.([][]byte); ok {
		hashes := make([]types.Hash, len(byteSlices))
		for i, byteArray := range byteSlices {
			if len(byteArray) != 32 {
				return nil, fmt.Errorf("item %d does not have length 32, length is %d", i, len(byteArray))
			}
			copy(hashes[i][:], byteArray)
		}
		return hashes, nil
	}

	// 尝试将 interface{} 转换为 []interface{}
	if slice, ok := data.([]interface{}); ok {
		// 将 []interface{} 转换为 [][]byte
		byteSlices := make([][]byte, len(slice))
		for i, item := range slice {
			if byteArray, ok := item.([]byte); ok {
				if len(byteArray) != 32 {
					return nil, fmt.Errorf("item %d does not have length 32, length is %d", i, len(byteArray))
				}
				byteSlices[i] = byteArray
			} else {
				return nil, fmt.Errorf("item %d is not of type []byte", i)
			}
		}
		// 转换为 []types.Hash
		hashes := make([]types.Hash, len(byteSlices))
		for i, byteArray := range byteSlices {
			copy(hashes[i][:], byteArray)
		}
		return hashes, nil
	}

	// 提供更详细的错误信息
	return nil, fmt.Errorf("interface{} is not of type [][]byte or []interface{}, got %T", data)
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
