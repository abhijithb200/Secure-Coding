import React, { useState, useEffect } from "react";
import "./Details.css";
import db from "./firebase";
import { doc, getDoc } from "firebase/firestore";

function Details() {
  const [data, setData] = useState([]);

  const myFunction = async () => {
    const docRef = doc(db, "vulnchart", "1");
    const docSnap = await getDoc(docRef);

    if (docSnap.exists()) {
      setData(docSnap.data());
      console.log("Document data:", docSnap.data());
    } else {
      // docSnap.data() will be undefined in this case
      console.log("No such document!");
    }
  };

  useEffect(() => {
    myFunction();
  }, []);
  return (
    <div className="Details_all">
      <h2 className="App_heading">CodeGuardian</h2>
      <div className="Details">
        <div className="Details_head">
          <p className="Details_headtitle">abhijithb200 : php-project</p>
          <p className="Details_headtime">{data.time}</p>
        </div>
        <div className="Details_SAST">
          <div className="Details_SASThead">
            <p className="Details_SASTheadhead">SAST</p>
            <p className="Details_SASTheadvuln">
              <span style={{ fontSize: "22px" }}>{data.SASTlength}</span>{" "}
              Vulnerabilities Found
            </p>
          </div>
          <div className="Details_SASTreport">
            <p className="Details_SASTreporthead">CodeGuardian output</p>
            <p
              className="Details_SASTreportbody"
              dangerouslySetInnerHTML={{ __html: data.SASTeverything }}
            ></p>
          </div>
        </div>
        <div className="Details_DAST">
          <p className="Details_SASTheadhead">DAST</p>
          <p className="Details_SASTheadvuln">
            <span style={{ fontSize: "22px" }}>{data.DASTlength}</span>{" "}
            Vulnerabilities Found
          </p>
        </div>
        <div className="Details_SQLi">
          <p className="Details_SASTreporthead">SQL Injection</p>
          <div className="Details_SASTreport">
            <p className="Details_SASTreporthead">SQLmap output</p>
            <p
              className="Details_SASTreportbody Details_SQLreportbody"
              dangerouslySetInnerHTML={{ __html: data.DASTsqlmapresult }}
            ></p>
          </div>
        </div>
        <div className="Details_SQLi Details_last">
          <p className="Details_SASTreporthead">
            Reflective Cross Site Scripting
          </p>
          <div className="Details_SASTreport">
            <p className="Details_SASTreporthead">PoC URL</p>
            <p
              className="Details_SASTreportbody"
              dangerouslySetInnerHTML={{ __html: data.DASTurl }}
            ></p>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Details;
