require! <[ fs cheerio request async ]>
{lists-to-obj, flatten} = require \prelude-ls

columns = <[ town village name ]>

tasks = for let i in [1 to 12]
  (cb) ->
    d = []
    err, res, body <- request.get "http://axe-level-1.herokuapp.com/lv2/?page=#{i}"
    $ = cheerio.load body
    row = $ 'table.table tr' .first!next!
    while row.text!
      d.push lists-to-obj columns, (row.children!map (,e) -> $(e).text!).to-array!
      row .= next!
    cb null d

err, data <- async.parallel tasks
data |> flatten |> JSON.stringify |> fs.write-file-sync 'output.json', _
