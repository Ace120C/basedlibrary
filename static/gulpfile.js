const {src, dest, watch, series} = require('gulp')
const sass = require('gulp-sass')(require('sass'))

function buildstyles() {
  return src('sass/**/*.sass')
  .pipe(sass())
  .pipe(dest('css'))
}

function watchtask(){
  watch(['sass/**/*.sass'], buildstyles)
}

exports.default = series(buildstyles, watchtask)
