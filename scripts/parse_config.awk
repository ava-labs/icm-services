/^\[profile.*\]$/ {
    # capture the current section-name
    section = substr($0, 10, length - 10)
}
section == profile && $1 == key {
    sub(/^"/,"", $3) # remove any leading double-quote
    sub(/^'/,"", $3) # remove any leading single-quote
    sub(/"$/, "", $3) # remove any trailing double-quote
    sub(/'$/, "", $3) # remove any trailing single-quote
    print $3
}
