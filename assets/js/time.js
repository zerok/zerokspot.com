document.querySelectorAll("article .dt-published").forEach((el) => {
  var postDate = new Date(el.getAttribute("datetime")).getTime();
  var now = new Date().getTime();
  var ageSeconds = now - postDate;
  var ageDays = (ageSeconds / (1000 * 60 * 60 * 24));
  el.innerHTML = el.innerHTML + "<span class=\"age\">(~" + Math.round(ageDays) + " day(s) ago)</span>";
});
