function merge(t1, t2) for k, v in pairs(t2) do table.insert(t1, v) end return t1 end

local function flatten (array)
  local ret = {}
  for _, v in ipairs(array) do
    if type(v) == 'table' then
      ret = merge(ret, flatten(v))
    else
      table.insert(ret, v)
    end
  end
  return ret
end

return flatten