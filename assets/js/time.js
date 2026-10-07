document.querySelectorAll("article .dt-published").forEach((el) => {
  var postDate = new Date(el.getAttribute("datetime")).getTime();
  var now = new Date().getTime();
  var ageSeconds = now - postDate;
  var ageDays = Math.round(ageSeconds / (1000 * 60 * 60 * 24));
  var ageYears = 0;
  if (ageDays > 365) {
    ageYears = Math.floor(ageDays / 365);
    ageDays = ageDays % 365;
  }
  var ageString = "";
  if (ageYears > 0) {
    ageString += ageYears + " ";
    if (ageYears != 1) {
      ageString += "years";
    } else {
      ageString += "year";
    }
  }
  if (ageDays > 0) {
    if (ageString != "") {
      ageString += " ";
    }
    ageString += ageDays + " ";
    if (ageDays != 1) {
      ageString += "days";
    } else {
      ageString += "day";
    }
  }
  el.innerHTML = el.innerHTML + "<span class=\"age\">(~" + ageString + " ago)</span>";
});
