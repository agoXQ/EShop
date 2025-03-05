-- 定义请求体内容
local json_body = '{"id": 1}'

-- request 函数定义了如何构造请求
function request()
    -- 使用 wrk.format 函数来构建请求，第一个参数是 HTTP 方法，第二个是 URL 路径，
    -- 第三个是表形式的 HTTP 头部，第四个是请求体。
    return wrk.format("POST", "/product/get", {["Content-Type"] = "application/json"}, json_body)
end

-- response 函数定义了如何处理响应（可选）
function response(status, headers, body)
    -- 如果你想检查非 200 状态码的响应，可以在这里添加逻辑
    if status ~= 200 then
        print("Non-200 response: " .. status)
    end
end