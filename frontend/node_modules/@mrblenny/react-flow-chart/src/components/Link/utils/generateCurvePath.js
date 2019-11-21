"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.generateCurvePath = function (startPos, endPos) {
    var width = Math.abs(startPos.x - endPos.x);
    var height = Math.abs(startPos.y - endPos.y);
    var leftToRight = startPos.x < endPos.x;
    var topToBottom = startPos.y < endPos.y;
    var isHorizontal = width > height;
    var start;
    var end;
    if (isHorizontal) {
        start = leftToRight ? startPos : endPos;
        end = leftToRight ? endPos : startPos;
    }
    else {
        start = topToBottom ? startPos : endPos;
        end = topToBottom ? endPos : startPos;
    }
    var curve = isHorizontal ? width / 3 : height / 3;
    var curveX = isHorizontal ? curve : 0;
    var curveY = isHorizontal ? 0 : curve;
    return "M" + start.x + "," + start.y + " C " + (start.x + curveX) + "," + (start.y + curveY) + " " + (end.x - curveX) + "," + (end.y - curveY) + " " + end.x + "," + end.y;
};
//# sourceMappingURL=generateCurvePath.js.map