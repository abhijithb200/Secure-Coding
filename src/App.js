import logo from "./logo.svg";
import "./App.css";
import db from "./firebase";
import {
  collection,
  query,
  where,
  onSnapshot,
  orderBy,
} from "firebase/firestore";

import React, { useState, useEffect } from "react";
import Details from "./Details";

function ResultCard({ time, SASTstatus, DASTstatus, DASTlength, SASTlength }) {
  return (
    <div className="ResultCard">
      <div className="ResultCard_top">
        <p className="ResultCard_head">abhijithb200 : php-project</p>
        <p className="ResultCard_headlast">{time}</p>
      </div>

      <div className="ResultCard_result">
        <div className="ResultCard_SAST">
          <p className="ResultCard_SASThead">SAST</p>
          <table>
            <tr>
              <td>Status</td>
              <td>
                :{" "}
                {SASTstatus && (
                  <span style={{ color: "green", fontWeight: "600" }}>
                    Fetched
                  </span>
                )}
              </td>
            </tr>
            <tr>
              <td>-</td>
            </tr>
            <tr>
              <td>Vulnerabilities</td>
              <td>
                :{" "}
                <span style={{ color: "red", fontWeight: "600" }}>
                  {SASTlength}
                </span>
              </td>
            </tr>
          </table>
        </div>
        <div className="ResultCard_DAST">
          <p className="ResultCard_DASThead">DAST</p>
          <table>
            <tr>
              <td>Status</td>
              <td>
                :{" "}
                {!DASTstatus ? (
                  <span style={{ color: "#00B5E6", fontWeight: "600" }}>
                    Queued
                  </span>
                ) : (
                  <span style={{ color: "green", fontWeight: "600" }}>
                    Fetched
                  </span>
                )}
              </td>
            </tr>
            <tr>
              <td>-</td>
            </tr>
            <tr>
              <td>Vulnerabilities</td>
              <td>
                :{" "}
                <span style={{ color: "red", fontWeight: "600" }}>
                  {DASTlength}
                </span>
              </td>
            </tr>
          </table>
        </div>
      </div>
    </div>
  );
}

function App() {
  const [data, setData] = useState([]);

  const myFunction = async () => {
    const q = query(collection(db, "vulnchart"), orderBy("time", "desc"));
    const unsubscribe = onSnapshot(q, (querySnapshot) => {
      setData(querySnapshot.docs.map((doc) => doc.data()));
    });

    console.log(data);
  };

  useEffect(() => {
    myFunction();
  }, []);

  return <Details />;
  // <div className="App">
  //   <h2 className="App_heading">CodeGuardian</h2>
  {
    /* <div
        className="content"
        dangerouslySetInnerHTML={{ __html: first }}
      ></div> */
  }

  {
    /* <div className="App_body">
        <p className="App_bodyhead">All Results</p>
        <div className="App_bodyarea">
          {data.map((dat) => (
            <ResultCard
              time={dat.time}
              SASTstatus={dat.SASTstatus}
              DASTstatus={dat.DASTstatus}
              SASTlength={dat.SASTlength}
              DASTlength={dat.DASTlength}
            />
          ))}
        </div>
      </div> */
  }
  <></>;
  // </div>
}

export default App;
