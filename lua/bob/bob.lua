bob = {}

function bob.hey(phrase)

    if phrase == nil or phrase == '' then
        return 'Fine, be that way.'
    elseif phrase:sub(#phrase, #phrase) == '?' then
        return 'Sure'
    elseif string.upper(phrase) == phrase then
        return 'Whoa, chill out!'
    else
        return 'Whatever'
    end
end

return bob