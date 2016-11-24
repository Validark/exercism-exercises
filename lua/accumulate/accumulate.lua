function accumulate(arr, func)
    mut = {}
    for i,v in ipairs(arr) do
        mut[i] = func(v)
    end
    return mut
end

return accumulate