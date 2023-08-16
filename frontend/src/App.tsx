import Sidebar from "./components/Sidebar";
import ConfirmDialog from "./components/confirmDialog";
import Flow from "./components/mindmap";
import MDEditore from "./components/mdEditor";

import { nanoid } from "nanoid/non-secure";

import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import { useEffect, useState } from "react";

import axiosInstance from "./utils/axiosInstance";
import { useForceUpdate } from "./hooks";
import useEditorStore from "./components/mdEditor/editorState";

type RootsResp = {
  id: string;
  label: string;
};

function App() {
  // toasts
  const notify = (msg: string) => toast.error(msg);
  // const notifySuccess = (msg: string) => toast.success(msg);
  const editorStates = useEditorStore();
  const [roots, setRoots] = useState<RootsResp[]>([]);
  const [rootsLoaded, setRootsLoaded] = useState<boolean>(false);
  const [selectedRoot, setSelectedRoot] = useState<string>("");
  const forceUpdate = useForceUpdate();

  // render categories API
  useEffect(() => {
    (async () => {
      try {
        const resp = await axiosInstance.get("/roots");
        if (resp.status === 200) {
          setRoots(resp.data);
          if (resp.data.length > 0) {
            setSelectedRoot(resp.data[0].id)
          }
        } else {
          notify(resp.data?.error);
        }
        setRootsLoaded(true);
      } catch (error) {
        setRootsLoaded(true);
        notify(String(error));
      }
    })();
  }, []);

  const handleOpen = (id: string) => {
    setSelectedRoot(id)
  };

  const handleAdd = async () => {
    const data: RootsResp = {
      id: nanoid(),
      label: "new title"
    }

    try {
      await axiosInstance.post("root", data);
      forceUpdate();
    } catch (error) {
      console.log(error);
      notify("there is an error in new title")
    }
  };

  const handleOpenEditor = (nodeId:string) => {
    // set editor Id
    // set isFlow to false
    console.log(nodeId)
  }
  handleOpenEditor("testtttt")

  return (
    <div className="app">
      <Sidebar onAdd={handleAdd}>
        <h2 className="text-3xl font-bold underline">Main Titles</h2>
        {!rootsLoaded && "loading"}
        {rootsLoaded && roots.length > 0 && (
          <div>
            {roots.map((el, idx) => (
              <button
                key={idx}
                onClick={() => handleOpen(el.id)}
                className="sidebutton"
              >
                {el.label}
                <div>deleted</div>
              </button>
            ))}
          </div>
        )}
      </Sidebar>
      <div className="main-content">
        {/* blank page */}
        {!editorStates.isOpen ? <Flow rootId={selectedRoot} /> : <MDEditore />}
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
