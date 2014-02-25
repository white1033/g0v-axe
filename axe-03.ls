require! <[fs cheerio request async]>
{lists-to-obj, flatten} = require \prelude-ls

columns = <[town village name]>
request = request.defaults jar: true

tasks = for let i from 1 to 76
  (cb) ->
    d = []
    url = | i == 1 => "http://axe-level-1.herokuapp.com/lv3/"
          | _ => "http://axe-level-1.herokuapp.com/lv3/?page=next"
    err, res, body <- request.get url
    $ = cheerio.load body
    row = $ 'table.table tr' .first!.next!
    while row.text!
      d.push lists-to-obj columns, (row.children!.map (,e) -> $(e).text!).to-array!
      row .= next!
    cb null, d

err, data <- async.series tasks
data |> flatten |> JSON.stringify |> fs.write-file-sync 'output.json', _
