window.onload = function() {
  var paperDom = document.getElementById("svg");
  var paper = Raphael("svg", 238, 175 + 2);

  var animationDelay = 400;

  var rings = paper.set();

  rings.push(paper.path("M135.879059,156.558065 C131.431506,152.485197 125.505977,150 119,150 C112.494023,150 106.568494,152.485197 102.120941,156.558065 L119,175 L135.879059,156.558065 Z"));

  rings.push(paper.path("M169.637167,119.674186 C156.294508,107.455587 138.517928,100 119,100 C99.4820676,100 81.7054827,107.455591 68.3628229,119.674195 L85.2418819,138.11613 C94.1369962,129.970391 105.988049,125 119,125 C132.011951,125 143.863004,129.970391 152.758109,138.116122 L169.637177,119.674195 Z"));

  rings.push(paper.path("M203.395295,82.7903257 C181.157529,62.4259849 151.529887,50 119,50 C86.4701143,50 56.8424743,62.4259836 34.6047084,82.7903225 L51.4837639,101.232261 C69.2739769,84.9407879 92.9760901,75 119,75 C145.02391,75 168.726023,84.9407879 186.516236,101.232261 L203.395295,82.7903257 Z"));

  rings.push(paper.path("M237.153413,45.906456 C206.02054,17.3963789 164.541842,0 119,0 C73.4581579,0 31.9794602,17.3963787 0.846587402,45.9064554 L17.7256458,64.3483908 C44.4109654,39.9111819 79.9641351,25 119,25 C158.035865,25 193.589035,39.9111819 220.274354,64.3483908 L237.153413,45.906456 Z"));

  rings.translate(0, 2);

  if (paperDom.classList.contains("green")) {
    turnOn(rings);
  } else if (paperDom.classList.contains("yellow")) {
    turnOff(rings[3]);
    turnOff(rings[2]);
    turnOn(rings[1]);
    turnOn(rings[0]);
  } else if (paperDom.classList.contains("red")) {
    turnOff(rings);

    window.setTimeout(animateRings, animationDelay, true, 0);
  }

  function turnOn(element) {
    element.attr({fill: "black", stroke: 1});
  }

  function turnOff(element) {
    element.attr({fill: "#d5d5d5", stroke: 0});
  }

  function animateRings(up, i) {
    if (i > 0 && i <= rings.length) {
      if (up) {
        turnOff(rings[i - 1]);
      } else {
        turnOff(rings[i + 1]);
      }
    } else if (i == 0) {
      if (up) {
        turnOff(rings[rings.length - 1]);
      } else {
        turnOff(rings[1]);
      }
    }

    turnOn(rings[i]);

    var nextI;
    var nextUp = up;

    if (i == rings.length - 1) {
      nextUp = false;
    } else if (i == 0) {
      nextUp = true;
    }

    if (nextUp) {
      nextI = (i + 1) % rings.length;
    } else {
      nextI = i - 1;
    }

    window.setTimeout(animateRings, animationDelay, nextUp, nextI);
  }
}