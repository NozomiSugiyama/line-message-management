declare const process: any;
import React from "react";
import ReactDOM from "react-dom";
import App from "src/App";

console.log(
    `
%c\n  (＼_(＼ \t
  （ =･ω･）\t
.c(,_ｕｕﾉ\t
%c%c
created by: https://github.com/NozomiSugiyama/line-message-management
NODE_ENV  : ${process.env.NODE_ENV}
VERSION   : ${process.env.VERSION}`,
    "color:#27365d;font-size:24px;line-height:24px;",
    "",
    "color:#bbb;"
);
ReactDOM.render(<App/>, document.getElementById("root"));
