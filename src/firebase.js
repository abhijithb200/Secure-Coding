import { initializeApp } from "firebase/app";
import { getFirestore } from "firebase/firestore";

// TODO: Replace the following with your app's Firebase project configuration
const firebaseConfig = {
  apiKey: "AIzaSyBWs7Pr2jkSjQsn1YdBelbeAK_s2ZhjUfU",
  authDomain: "codeguardian-1035d.firebaseapp.com",
  projectId: "codeguardian-1035d",
  storageBucket: "codeguardian-1035d.appspot.com",
  messagingSenderId: "619280214072",
  appId: "1:619280214072:web:fd1ca091e0a12d82e21ecd",
};

const app = initializeApp(firebaseConfig);
const db = getFirestore(app);

export default db;
