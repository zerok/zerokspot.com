const Gulp = require('gulp');
const Sass = require('gulp-sass');
const rename = require('gulp-rename');

function mainCSS() {
    return Gulp.src('static/sass/main.scss')
        .pipe(Sass())
        .pipe(Gulp.dest('static/css'));
}

function previewCSS() {
    return Gulp.src('static/sass/preview.scss')
        .pipe(Sass())
        .pipe(Gulp.dest('static/css'));
}

exports.copyNormalized = function copyNormalized() {
    return Gulp.src('node_modules/normalize.css/normalize.css')
        .pipe(rename({basename: '_normalize', extname: '.scss'}))
        .pipe(Gulp.dest('static/sass'));
};

exports.sass = Gulp.series(exports.copyNormalized, Gulp.parallel(mainCSS, previewCSS));

exports.sassWatch = function sassWatch() {
    return Gulp.watch('static/sass/*.scss', Gulp.parallel(mainCSS, previewCSS));
}
exports.default = Gulp.series(exports.sass, exports.sassWatch);
