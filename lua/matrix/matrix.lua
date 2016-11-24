-- initialize
local function parse_matrix(str_matrix)
  rows = {}
  for c in str_matrix:gmatch('[%d ]+') do
    row = {}
    for r in c:gmatch('%d+') do
      table.insert(row, tonumber(r))
    end
    table.insert(rows, row)
  end
  return rows
end

return function(str_matrix)

  matrix =  parse_matrix(str_matrix)

  local function row(r)
      return matrix[r]
  end

  local function column(col)
    ret = {}
    for c, r in ipairs(matrix) do
      table.insert(ret, r[col])
    end
    return ret
  end

  return {
    row=row,
    column=column
  }
end