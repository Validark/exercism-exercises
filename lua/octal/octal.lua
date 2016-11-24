local function to_decimal(octal)
    local sum = 0
    for i=#octal, 1, -1 do
       local val = tonumber(octal:sub(i,i))
       if val == nil or val > 7 then return 0 end
       sum = sum + val * 8^(#octal-i)
    end
    return sum
end

return function(octal)
   return {to_decimal = function () return to_decimal(octal) end}
end