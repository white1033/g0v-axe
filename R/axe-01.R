library(XML)
library(jsonlite)
html.data <- readHTMLTable('http://axe-level-1.herokuapp.com/', colClasses = c('character', 'integer', 'integer', 'integer', 'integer', 'integer'), stringsAsFactors = F)[[1]]
colnames(html.data)[1] <- 'name'
html.data$grades <- as.data.frame(html.data[-1])
write(toJSON(html.data[, c('name', 'grades')]), 'output.json')
