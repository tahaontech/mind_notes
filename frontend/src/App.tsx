import Sidebar from "./components/Sidebar";
import ConfirmDialog from "./components/confirmDialog";
import Flow from "./components/mindmap";

import { ToastContainer } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

function App() {
  return (
    <div className="app">
      <Sidebar>
        <h2>Main Titles</h2>
        <button className="sidebutton">Philosophy</button>
      </Sidebar>
      <div className="main-content">
        <Flow />
        {/* <h1>markdown editor selectedId</h1> */}
      </div>
      <ConfirmDialog />
        <ToastContainer
          position="bottom-center"
          autoClose={5000}
          hideProgressBar={false}
          newestOnTop={false}
          closeOnClick
          rtl={false}
          pauseOnFocusLoss
          draggable
          pauseOnHover
          theme="light"
        />
    </div>
  );
}

export default App;
