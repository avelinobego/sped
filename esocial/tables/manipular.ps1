$arqivos = Get-ChildItem -Name -Include *.csv, *.txt

foreach ($nome in $arqivos){
    $temp = [io.path]::ChangeExtension([Text.Encoding]::ASCII.GetString([Text.Encoding]::GetEncoding("Cyrillic").GetBytes($nome)), "txt")
    Write-Output "$temp"
    Rename-Item -Path $nome -NewName $temp
}