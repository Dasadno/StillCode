package runner

import (
	"fmt"
	"strings"
)

// Language templates for LeetCode-style code execution.
// Each template wraps user code with a harness that:
// 1. Parses JSON input from stdin
// 2. Calls the user's function with parsed arguments
// 3. Prints the result as JSON to stdout

// Template placeholders:
// {USER_CODE} - The user's solution function
// {FUNCTION_NAME} - The name of the function to call
// {ARGS} - The arguments to pass to the function (e.g., "data["nums"], data["target"]")

const pythonTemplate = `import json
import sys
from typing import List, Optional, Dict, Set, Tuple, Any
from collections import defaultdict, Counter, deque, OrderedDict
from heapq import heappush, heappop, heapify, heappushpop, nlargest, nsmallest
from bisect import bisect_left, bisect_right, insort_left, insort_right
from functools import lru_cache, reduce
from itertools import permutations, combinations, product, accumulate
import math
import re
import string

# User's solution
{USER_CODE}

# Parse input and execute
if __name__ == "__main__":
    data = json.loads(sys.stdin.read())
    result = {FUNCTION_NAME}({ARGS})
    print(json.dumps(result))
`

const javascriptTemplate = `const readline = require('readline');

// User's solution
{USER_CODE}

// Parse input and execute
let input = '';
process.stdin.setEncoding('utf8');
process.stdin.on('data', chunk => input += chunk);
process.stdin.on('end', () => {
    try {
        const data = JSON.parse(input);
        const result = {FUNCTION_NAME}({ARGS});
        console.log(JSON.stringify(result));
    } catch (e) {
        console.error(e.message);
        process.exit(1);
    }
});
`

const goTemplate = `package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// Suppress unused import errors
var _ = math.MaxInt
var _ = sort.Ints
var _ = strconv.Itoa
var _ = strings.Contains
var _ = unicode.IsLetter

// User's solution
{USER_CODE}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var input []byte
	for {
		line, err := reader.ReadBytes('\n')
		input = append(input, line...)
		if err != nil {
			break
		}
	}

	var data map[string]interface{}
	if err := json.Unmarshal(input, &data); err != nil {
		fmt.Fprintf(os.Stderr, "JSON parse error: %v\n", err)
		os.Exit(1)
	}

	// Convert data to typed arguments
	{ARG_CONVERSIONS}

	result := {FUNCTION_NAME}({ARGS})

	output, err := json.Marshal(result)
	if err != nil {
		fmt.Fprintf(os.Stderr, "JSON marshal error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(output))
}
`

const javaTemplate = `import java.io.*;
import java.util.*;
import com.google.gson.*;
import com.google.gson.reflect.*;

public class Main {
    // User's solution
    {USER_CODE}

    public static void main(String[] args) {
        try {
            BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));
            StringBuilder sb = new StringBuilder();
            String line;
            while ((line = reader.readLine()) != null) {
                sb.append(line);
            }

            Gson gson = new Gson();
            JsonObject data = JsonParser.parseString(sb.toString()).getAsJsonObject();

            Solution solution = new Solution();
            {ARG_PARSING}
            {RESULT_TYPE} result = solution.{FUNCTION_NAME}({ARGS});

            System.out.println(gson.toJson(result));
        } catch (Exception e) {
            System.err.println(e.getMessage());
            System.exit(1);
        }
    }
}
`

// Alternative Java template without external dependencies (using simple JSON parsing)
const javaTemplateSimple = `import java.io.*;
import java.util.*;
import java.util.regex.*;

public class Main {
    // User's solution
    {USER_CODE}

    // Simple JSON parser for common types
    static class SimpleJson {
        private String json;
        private int pos;

        public SimpleJson(String json) {
            this.json = json.trim();
            this.pos = 0;
        }

        public Map<String, Object> parseObject() {
            Map<String, Object> map = new HashMap<>();
            expect('{');
            skipWhitespace();
            if (peek() != '}') {
                do {
                    skipWhitespace();
                    String key = parseString();
                    skipWhitespace();
                    expect(':');
                    skipWhitespace();
                    Object value = parseValue();
                    map.put(key, value);
                    skipWhitespace();
                } while (match(','));
            }
            expect('}');
            return map;
        }

        private Object parseValue() {
            skipWhitespace();
            char c = peek();
            if (c == '"') return parseString();
            if (c == '[') return parseArray();
            if (c == '{') return parseObject();
            if (c == 't' || c == 'f') return parseBoolean();
            if (c == 'n') return parseNull();
            return parseNumber();
        }

        private String parseString() {
            expect('"');
            StringBuilder sb = new StringBuilder();
            while (peek() != '"') {
                char c = advance();
                if (c == '\\') {
                    char esc = advance();
                    switch (esc) {
                        case 'n': sb.append('\n'); break;
                        case 't': sb.append('\t'); break;
                        case 'r': sb.append('\r'); break;
                        case '"': sb.append('"'); break;
                        case '\\': sb.append('\\'); break;
                        default: sb.append(esc);
                    }
                } else {
                    sb.append(c);
                }
            }
            expect('"');
            return sb.toString();
        }

        private List<Object> parseArray() {
            List<Object> list = new ArrayList<>();
            expect('[');
            skipWhitespace();
            if (peek() != ']') {
                do {
                    skipWhitespace();
                    list.add(parseValue());
                    skipWhitespace();
                } while (match(','));
            }
            expect(']');
            return list;
        }

        private Number parseNumber() {
            int start = pos;
            if (peek() == '-') advance();
            while (pos < json.length() && (Character.isDigit(peek()) || peek() == '.' || peek() == 'e' || peek() == 'E' || peek() == '+' || peek() == '-')) {
                if (peek() == '.' || peek() == 'e' || peek() == 'E') {
                    advance();
                    while (pos < json.length() && (Character.isDigit(peek()) || peek() == '+' || peek() == '-')) advance();
                    return Double.parseDouble(json.substring(start, pos));
                }
                advance();
            }
            return Long.parseLong(json.substring(start, pos));
        }

        private Boolean parseBoolean() {
            if (json.substring(pos).startsWith("true")) { pos += 4; return true; }
            if (json.substring(pos).startsWith("false")) { pos += 5; return false; }
            throw new RuntimeException("Invalid boolean at " + pos);
        }

        private Object parseNull() {
            if (json.substring(pos).startsWith("null")) { pos += 4; return null; }
            throw new RuntimeException("Invalid null at " + pos);
        }

        private void skipWhitespace() { while (pos < json.length() && Character.isWhitespace(json.charAt(pos))) pos++; }
        private char peek() { return pos < json.length() ? json.charAt(pos) : '\0'; }
        private char advance() { return json.charAt(pos++); }
        private void expect(char c) { if (advance() != c) throw new RuntimeException("Expected '" + c + "' at " + (pos-1)); }
        private boolean match(char c) { if (peek() == c) { advance(); return true; } return false; }

        // Conversion helpers
        public static int[] toIntArray(List<Object> list) {
            int[] arr = new int[list.size()];
            for (int i = 0; i < list.size(); i++) arr[i] = ((Number) list.get(i)).intValue();
            return arr;
        }

        public static int[][] toInt2DArray(List<Object> list) {
            int[][] arr = new int[list.size()][];
            for (int i = 0; i < list.size(); i++) arr[i] = toIntArray((List<Object>) list.get(i));
            return arr;
        }

        public static String[] toStringArray(List<Object> list) {
            String[] arr = new String[list.size()];
            for (int i = 0; i < list.size(); i++) arr[i] = (String) list.get(i);
            return arr;
        }

        public static List<Integer> toIntegerList(List<Object> list) {
            List<Integer> result = new ArrayList<>();
            for (Object o : list) result.add(((Number) o).intValue());
            return result;
        }

        public static List<List<Integer>> toInteger2DList(List<Object> list) {
            List<List<Integer>> result = new ArrayList<>();
            for (Object o : list) result.add(toIntegerList((List<Object>) o));
            return result;
        }

        // JSON output helpers
        public static String toJson(Object obj) {
            if (obj == null) return "null";
            if (obj instanceof Boolean) return obj.toString();
            if (obj instanceof Number) return obj.toString();
            if (obj instanceof String) return "\"" + escapeString((String) obj) + "\"";
            if (obj instanceof int[]) {
                int[] arr = (int[]) obj;
                StringBuilder sb = new StringBuilder("[");
                for (int i = 0; i < arr.length; i++) {
                    if (i > 0) sb.append(",");
                    sb.append(arr[i]);
                }
                return sb.append("]").toString();
            }
            if (obj instanceof int[][]) {
                int[][] arr = (int[][]) obj;
                StringBuilder sb = new StringBuilder("[");
                for (int i = 0; i < arr.length; i++) {
                    if (i > 0) sb.append(",");
                    sb.append(toJson(arr[i]));
                }
                return sb.append("]").toString();
            }
            if (obj instanceof String[]) {
                String[] arr = (String[]) obj;
                StringBuilder sb = new StringBuilder("[");
                for (int i = 0; i < arr.length; i++) {
                    if (i > 0) sb.append(",");
                    sb.append(toJson(arr[i]));
                }
                return sb.append("]").toString();
            }
            if (obj instanceof List) {
                List<?> list = (List<?>) obj;
                StringBuilder sb = new StringBuilder("[");
                for (int i = 0; i < list.size(); i++) {
                    if (i > 0) sb.append(",");
                    sb.append(toJson(list.get(i)));
                }
                return sb.append("]").toString();
            }
            return obj.toString();
        }

        private static String escapeString(String s) {
            return s.replace("\\", "\\\\").replace("\"", "\\\"").replace("\n", "\\n").replace("\r", "\\r").replace("\t", "\\t");
        }
    }

    public static void main(String[] args) {
        try {
            BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));
            StringBuilder sb = new StringBuilder();
            String line;
            while ((line = reader.readLine()) != null) {
                sb.append(line);
            }

            SimpleJson parser = new SimpleJson(sb.toString());
            Map<String, Object> data = parser.parseObject();

            Solution solution = new Solution();
            {ARG_PARSING}
            {RESULT_TYPE} result = solution.{FUNCTION_NAME}({ARGS});

            System.out.println(SimpleJson.toJson(result));
        } catch (Exception e) {
            e.printStackTrace();
            System.exit(1);
        }
    }
}
`

const cppTemplate = `#include <iostream>
#include <sstream>
#include <string>
#include <vector>
#include <map>
#include <algorithm>
#include <cmath>
#include <climits>
#include <unordered_map>
#include <unordered_set>
#include <queue>
#include <stack>
#include <set>

using namespace std;

// Simple JSON parser for competitive programming
class JSON {
public:
    static string trim(const string& s) {
        size_t start = s.find_first_not_of(" \t\n\r");
        if (start == string::npos) return "";
        size_t end = s.find_last_not_of(" \t\n\r");
        return s.substr(start, end - start + 1);
    }

    static pair<string, size_t> parseString(const string& s, size_t pos) {
        if (s[pos] != '"') throw runtime_error("Expected string");
        pos++;
        string result;
        while (pos < s.length() && s[pos] != '"') {
            if (s[pos] == '\\' && pos + 1 < s.length()) {
                pos++;
                switch (s[pos]) {
                    case 'n': result += '\n'; break;
                    case 't': result += '\t'; break;
                    case 'r': result += '\r'; break;
                    case '"': result += '"'; break;
                    case '\\': result += '\\'; break;
                    default: result += s[pos];
                }
            } else {
                result += s[pos];
            }
            pos++;
        }
        return {result, pos + 1};
    }

    static pair<long long, size_t> parseInt(const string& s, size_t pos) {
        size_t start = pos;
        if (s[pos] == '-') pos++;
        while (pos < s.length() && isdigit(s[pos])) pos++;
        return {stoll(s.substr(start, pos - start)), pos};
    }

    static pair<double, size_t> parseDouble(const string& s, size_t pos) {
        size_t start = pos;
        if (s[pos] == '-') pos++;
        while (pos < s.length() && (isdigit(s[pos]) || s[pos] == '.' || s[pos] == 'e' || s[pos] == 'E' || s[pos] == '+' || s[pos] == '-')) pos++;
        return {stod(s.substr(start, pos - start)), pos};
    }

    static pair<bool, size_t> parseBool(const string& s, size_t pos) {
        if (s.substr(pos, 4) == "true") return {true, pos + 4};
        if (s.substr(pos, 5) == "false") return {false, pos + 5};
        throw runtime_error("Expected boolean");
    }

    static size_t skipWhitespace(const string& s, size_t pos) {
        while (pos < s.length() && isspace(s[pos])) pos++;
        return pos;
    }

    static pair<vector<int>, size_t> parseIntArray(const string& s, size_t pos) {
        vector<int> result;
        pos = skipWhitespace(s, pos);
        if (s[pos] != '[') throw runtime_error("Expected array");
        pos = skipWhitespace(s, pos + 1);
        if (s[pos] != ']') {
            while (true) {
                pos = skipWhitespace(s, pos);
                auto [val, newPos] = parseInt(s, pos);
                result.push_back((int)val);
                pos = skipWhitespace(s, newPos);
                if (s[pos] == ']') break;
                if (s[pos] != ',') throw runtime_error("Expected ',' or ']'");
                pos++;
            }
        }
        return {result, pos + 1};
    }

    static pair<vector<vector<int>>, size_t> parseInt2DArray(const string& s, size_t pos) {
        vector<vector<int>> result;
        pos = skipWhitespace(s, pos);
        if (s[pos] != '[') throw runtime_error("Expected 2D array");
        pos = skipWhitespace(s, pos + 1);
        if (s[pos] != ']') {
            while (true) {
                pos = skipWhitespace(s, pos);
                auto [arr, newPos] = parseIntArray(s, pos);
                result.push_back(arr);
                pos = skipWhitespace(s, newPos);
                if (s[pos] == ']') break;
                if (s[pos] != ',') throw runtime_error("Expected ',' or ']'");
                pos++;
            }
        }
        return {result, pos + 1};
    }

    static pair<vector<string>, size_t> parseStringArray(const string& s, size_t pos) {
        vector<string> result;
        pos = skipWhitespace(s, pos);
        if (s[pos] != '[') throw runtime_error("Expected array");
        pos = skipWhitespace(s, pos + 1);
        if (s[pos] != ']') {
            while (true) {
                pos = skipWhitespace(s, pos);
                auto [str, newPos] = parseString(s, pos);
                result.push_back(str);
                pos = skipWhitespace(s, newPos);
                if (s[pos] == ']') break;
                if (s[pos] != ',') throw runtime_error("Expected ',' or ']'");
                pos++;
            }
        }
        return {result, pos + 1};
    }

    static map<string, size_t> parseObjectKeys(const string& s) {
        map<string, size_t> keys;
        size_t pos = skipWhitespace(s, 0);
        if (s[pos] != '{') throw runtime_error("Expected object");
        pos = skipWhitespace(s, pos + 1);
        if (s[pos] != '}') {
            while (true) {
                pos = skipWhitespace(s, pos);
                auto [key, newPos] = parseString(s, pos);
                pos = skipWhitespace(s, newPos);
                if (s[pos] != ':') throw runtime_error("Expected ':'");
                pos = skipWhitespace(s, pos + 1);
                keys[key] = pos;
                // Skip the value
                pos = skipValue(s, pos);
                pos = skipWhitespace(s, pos);
                if (s[pos] == '}') break;
                if (s[pos] != ',') throw runtime_error("Expected ',' or '}'");
                pos++;
            }
        }
        return keys;
    }

    static size_t skipValue(const string& s, size_t pos) {
        pos = skipWhitespace(s, pos);
        if (s[pos] == '"') {
            auto [_, newPos] = parseString(s, pos);
            return newPos;
        }
        if (s[pos] == '[') {
            pos++;
            int depth = 1;
            while (depth > 0 && pos < s.length()) {
                if (s[pos] == '[') depth++;
                else if (s[pos] == ']') depth--;
                else if (s[pos] == '"') {
                    auto [_, newPos] = parseString(s, pos);
                    pos = newPos;
                    continue;
                }
                pos++;
            }
            return pos;
        }
        if (s[pos] == '{') {
            pos++;
            int depth = 1;
            while (depth > 0 && pos < s.length()) {
                if (s[pos] == '{') depth++;
                else if (s[pos] == '}') depth--;
                else if (s[pos] == '"') {
                    auto [_, newPos] = parseString(s, pos);
                    pos = newPos;
                    continue;
                }
                pos++;
            }
            return pos;
        }
        if (s.substr(pos, 4) == "true") return pos + 4;
        if (s.substr(pos, 5) == "false") return pos + 5;
        if (s.substr(pos, 4) == "null") return pos + 4;
        // Number
        if (s[pos] == '-') pos++;
        while (pos < s.length() && (isdigit(s[pos]) || s[pos] == '.' || s[pos] == 'e' || s[pos] == 'E' || s[pos] == '+' || s[pos] == '-')) pos++;
        return pos;
    }

    // JSON output functions
    static string toJson(int val) { return to_string(val); }
    static string toJson(long long val) { return to_string(val); }
    static string toJson(double val) { return to_string(val); }
    static string toJson(bool val) { return val ? "true" : "false"; }
    static string toJson(const string& val) {
        string result = "\"";
        for (char c : val) {
            switch (c) {
                case '"': result += "\\\""; break;
                case '\\': result += "\\\\"; break;
                case '\n': result += "\\n"; break;
                case '\r': result += "\\r"; break;
                case '\t': result += "\\t"; break;
                default: result += c;
            }
        }
        return result + "\"";
    }
    static string toJson(const vector<int>& arr) {
        string result = "[";
        for (size_t i = 0; i < arr.size(); i++) {
            if (i > 0) result += ",";
            result += to_string(arr[i]);
        }
        return result + "]";
    }
    static string toJson(const vector<long long>& arr) {
        string result = "[";
        for (size_t i = 0; i < arr.size(); i++) {
            if (i > 0) result += ",";
            result += to_string(arr[i]);
        }
        return result + "]";
    }
    static string toJson(const vector<string>& arr) {
        string result = "[";
        for (size_t i = 0; i < arr.size(); i++) {
            if (i > 0) result += ",";
            result += toJson(arr[i]);
        }
        return result + "]";
    }
    static string toJson(const vector<bool>& arr) {
        string result = "[";
        for (size_t i = 0; i < arr.size(); i++) {
            if (i > 0) result += ",";
            result += arr[i] ? "true" : "false";
        }
        return result + "]";
    }
    static string toJson(const vector<vector<int>>& arr) {
        string result = "[";
        for (size_t i = 0; i < arr.size(); i++) {
            if (i > 0) result += ",";
            result += toJson(arr[i]);
        }
        return result + "]";
    }
    static string toJson(const vector<vector<string>>& arr) {
        string result = "[";
        for (size_t i = 0; i < arr.size(); i++) {
            if (i > 0) result += ",";
            result += toJson(arr[i]);
        }
        return result + "]";
    }
};

// User's solution
{USER_CODE}

int main() {
    ios_base::sync_with_stdio(false);
    cin.tie(nullptr);

    string input;
    string line;
    while (getline(cin, line)) {
        input += line;
    }

    try {
        auto keys = JSON::parseObjectKeys(input);
        {ARG_PARSING}

        Solution solution;
        auto result = solution.{FUNCTION_NAME}({ARGS});
        cout << JSON::toJson(result) << endl;
    } catch (const exception& e) {
        cerr << "Error: " << e.what() << endl;
        return 1;
    }

    return 0;
}
`

// InputParam represents a parameter in the function signature
type InputParam struct {
	Name string `json:"name"` // Parameter name (e.g., "nums", "target")
	Type string `json:"type"` // Parameter type (e.g., "[]int", "int", "string", "[][]int")
}

// WrapUserCode wraps user code with the appropriate language template
// language: the programming language (python, javascript, go, java, cpp)
// userCode: the user's solution code
// functionName: the name of the function to call (e.g., "twoSum")
// params: the parameters to pass to the function with their types
func WrapUserCode(language, userCode, functionName string, params []InputParam) string {
	lang := strings.ToLower(language)

	switch lang {
	case "python", "py":
		return wrapPython(userCode, functionName, params)
	case "javascript", "js":
		return wrapJavaScript(userCode, functionName, params)
	case "go", "golang":
		return wrapGo(userCode, functionName, params)
	case "java":
		return wrapJava(userCode, functionName, params)
	case "cpp", "c++":
		return wrapCpp(userCode, functionName, params)
	default:
		return userCode // Return unchanged if language not supported
	}
}

func wrapPython(userCode, functionName string, params []InputParam) string {
	// Build argument list: data["nums"], data["target"]
	var args []string
	for _, p := range params {
		args = append(args, fmt.Sprintf(`data["%s"]`, p.Name))
	}

	code := pythonTemplate
	code = strings.Replace(code, "{USER_CODE}", userCode, 1)
	code = strings.Replace(code, "{FUNCTION_NAME}", functionName, 1)
	code = strings.Replace(code, "{ARGS}", strings.Join(args, ", "), 1)

	return code
}

func wrapJavaScript(userCode, functionName string, params []InputParam) string {
	// Build argument list: data.nums, data.target
	var args []string
	for _, p := range params {
		args = append(args, fmt.Sprintf("data.%s", p.Name))
	}

	code := javascriptTemplate
	code = strings.Replace(code, "{USER_CODE}", userCode, 1)
	code = strings.Replace(code, "{FUNCTION_NAME}", functionName, 1)
	code = strings.Replace(code, "{ARGS}", strings.Join(args, ", "), 1)

	return code
}

func wrapGo(userCode, functionName string, params []InputParam) string {
	// Build type conversions and argument list for Go
	var conversions []string
	var args []string

	for _, p := range params {
		varName := p.Name
		conversion := generateGoConversion(varName, p.Type)
		conversions = append(conversions, conversion)
		args = append(args, varName)
	}

	code := goTemplate
	code = strings.Replace(code, "{USER_CODE}", userCode, 1)
	code = strings.Replace(code, "{FUNCTION_NAME}", functionName, 1)
	code = strings.Replace(code, "{ARG_CONVERSIONS}", strings.Join(conversions, "\n\t"), 1)
	code = strings.Replace(code, "{ARGS}", strings.Join(args, ", "), 1)

	return code
}

func generateGoConversion(name, typ string) string {
	switch typ {
	case "int", "integer":
		return fmt.Sprintf(`%s := int(data["%s"].(float64))`, name, name)
	case "int64", "long":
		return fmt.Sprintf(`%s := int64(data["%s"].(float64))`, name, name)
	case "float", "float64", "double":
		return fmt.Sprintf(`%s := data["%s"].(float64)`, name, name)
	case "string":
		return fmt.Sprintf(`%s := data["%s"].(string)`, name, name)
	case "bool", "boolean":
		return fmt.Sprintf(`%s := data["%s"].(bool)`, name, name)
	case "[]int", "[]integer":
		return fmt.Sprintf(`%sRaw := data["%s"].([]interface{})
	%s := make([]int, len(%sRaw))
	for i, v := range %sRaw {
		%s[i] = int(v.(float64))
	}`, name, name, name, name, name, name)
	case "[]string":
		return fmt.Sprintf(`%sRaw := data["%s"].([]interface{})
	%s := make([]string, len(%sRaw))
	for i, v := range %sRaw {
		%s[i] = v.(string)
	}`, name, name, name, name, name, name)
	case "[][]int":
		return fmt.Sprintf(`%sRaw := data["%s"].([]interface{})
	%s := make([][]int, len(%sRaw))
	for i, row := range %sRaw {
		rowData := row.([]interface{})
		%s[i] = make([]int, len(rowData))
		for j, v := range rowData {
			%s[i][j] = int(v.(float64))
		}
	}`, name, name, name, name, name, name, name)
	case "[][]string":
		return fmt.Sprintf(`%sRaw := data["%s"].([]interface{})
	%s := make([][]string, len(%sRaw))
	for i, row := range %sRaw {
		rowData := row.([]interface{})
		%s[i] = make([]string, len(rowData))
		for j, v := range rowData {
			%s[i][j] = v.(string)
		}
	}`, name, name, name, name, name, name, name)
	default:
		// Default: try to use as interface{}
		return fmt.Sprintf(`%s := data["%s"]`, name, name)
	}
}

func wrapJava(userCode, functionName string, params []InputParam) string {
	// Build type conversions and argument list for Java
	var parsing []string
	var args []string

	for _, p := range params {
		varName := p.Name
		javaType, conversion := generateJavaConversion(varName, p.Type)
		parsing = append(parsing, fmt.Sprintf("%s %s = %s;", javaType, varName, conversion))
		args = append(args, varName)
	}

	// Determine result type (simplified - would need more context in real impl)
	resultType := "Object"

	// Transform "class Solution" to "static class Solution" for inner class
	modifiedCode := strings.Replace(userCode, "class Solution", "static class Solution", 1)

	code := javaTemplateSimple
	code = strings.Replace(code, "{USER_CODE}", modifiedCode, 1)
	code = strings.Replace(code, "{FUNCTION_NAME}", functionName, 1)
	code = strings.Replace(code, "{ARG_PARSING}", strings.Join(parsing, "\n            "), 1)
	code = strings.Replace(code, "{RESULT_TYPE}", resultType, 1)
	code = strings.Replace(code, "{ARGS}", strings.Join(args, ", "), 1)

	return code
}

func generateJavaConversion(name, typ string) (javaType, conversion string) {
	switch typ {
	case "int", "integer":
		return "int", fmt.Sprintf(`((Number) data.get("%s")).intValue()`, name)
	case "long", "int64":
		return "long", fmt.Sprintf(`((Number) data.get("%s")).longValue()`, name)
	case "double", "float", "float64":
		return "double", fmt.Sprintf(`((Number) data.get("%s")).doubleValue()`, name)
	case "string":
		return "String", fmt.Sprintf(`(String) data.get("%s")`, name)
	case "bool", "boolean":
		return "boolean", fmt.Sprintf(`(Boolean) data.get("%s")`, name)
	case "[]int", "[]integer":
		return "int[]", fmt.Sprintf(`SimpleJson.toIntArray((List<Object>) data.get("%s"))`, name)
	case "[]string":
		return "String[]", fmt.Sprintf(`SimpleJson.toStringArray((List<Object>) data.get("%s"))`, name)
	case "[][]int":
		return "int[][]", fmt.Sprintf(`SimpleJson.toInt2DArray((List<Object>) data.get("%s"))`, name)
	case "List<Integer>":
		return "List<Integer>", fmt.Sprintf(`SimpleJson.toIntegerList((List<Object>) data.get("%s"))`, name)
	case "List<List<Integer>>":
		return "List<List<Integer>>", fmt.Sprintf(`SimpleJson.toInteger2DList((List<Object>) data.get("%s"))`, name)
	default:
		return "Object", fmt.Sprintf(`data.get("%s")`, name)
	}
}

func wrapCpp(userCode, functionName string, params []InputParam) string {
	// Build type conversions and argument list for C++
	var parsing []string
	var args []string

	for _, p := range params {
		varName := p.Name
		cppType, conversion := generateCppConversion(varName, p.Type)
		parsing = append(parsing, fmt.Sprintf("%s %s = %s;", cppType, varName, conversion))
		args = append(args, varName)
	}

	code := cppTemplate
	code = strings.Replace(code, "{USER_CODE}", userCode, 1)
	code = strings.Replace(code, "{FUNCTION_NAME}", functionName, 1)
	code = strings.Replace(code, "{ARG_PARSING}", strings.Join(parsing, "\n        "), 1)
	code = strings.Replace(code, "{ARGS}", strings.Join(args, ", "), 1)

	return code
}

func generateCppConversion(name, typ string) (cppType, conversion string) {
	switch typ {
	case "int", "integer":
		return "int", fmt.Sprintf(`(int)JSON::parseInt(input, keys["%s"]).first`, name)
	case "long", "int64":
		return "long long", fmt.Sprintf(`JSON::parseInt(input, keys["%s"]).first`, name)
	case "double", "float", "float64":
		return "double", fmt.Sprintf(`JSON::parseDouble(input, keys["%s"]).first`, name)
	case "string":
		return "string", fmt.Sprintf(`JSON::parseString(input, keys["%s"]).first`, name)
	case "bool", "boolean":
		return "bool", fmt.Sprintf(`JSON::parseBool(input, keys["%s"]).first`, name)
	case "[]int", "[]integer":
		return "vector<int>", fmt.Sprintf(`JSON::parseIntArray(input, keys["%s"]).first`, name)
	case "[]string":
		return "vector<string>", fmt.Sprintf(`JSON::parseStringArray(input, keys["%s"]).first`, name)
	case "[][]int":
		return "vector<vector<int>>", fmt.Sprintf(`JSON::parseInt2DArray(input, keys["%s"]).first`, name)
	default:
		return "auto", fmt.Sprintf(`/* unsupported type for %s */`, name)
	}
}

// ConvenienceWrapper provides a simplified interface for common problem types
// This is useful when you know the function signature ahead of time
func ConvenienceWrapper(language, userCode, functionName string, paramNames []string, paramTypes []string) string {
	if len(paramNames) != len(paramTypes) {
		return userCode // Return unchanged if mismatch
	}

	params := make([]InputParam, len(paramNames))
	for i := range paramNames {
		params[i] = InputParam{Name: paramNames[i], Type: paramTypes[i]}
	}

	return WrapUserCode(language, userCode, functionName, params)
}
