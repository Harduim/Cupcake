import '../styles/globals.css'
// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import { getAnalytics } from "firebase/analytics";
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
  apiKey: "AIzaSyDoiTV9gcrposaaqG1NDy9Fa5s5RDZORAU",
  authDomain: "cupcake-io.firebaseapp.com",
  databaseURL: "https://cupcake-io-default-rtdb.firebaseio.com",
  projectId: "cupcake-io",
  storageBucket: "cupcake-io.appspot.com",
  messagingSenderId: "676584086694",
  appId: "1:676584086694:web:2faa1c76610fddafb1e867",
  measurementId: "G-4NGFLF766Y"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
const analytics = getAnalytics(app);

function MyApp({ Component, pageProps }) {
  return <Component {...pageProps} />
}

export default MyApp
