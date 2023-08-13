import { useLayoutEffect, useEffect, useRef } from "react";
import { Handle, NodeProps, Position } from "reactflow";

import useStore from "../store";

import DragIcon from "../../Icons/DragIcon";
import DeleteIcon from "../../Icons/DeleteIcon";
// import BookIcon from "../../Icons/BookIcon";
import useConfirmDialogStore from "../../confirmDialog/confirmDialogState";

import { toast } from "react-toastify";
import BookIcon from "../../Icons/BookIcon";

export type NodeData = {
  label: string;
  root: boolean;
};

function MindMapNode({ id, data }: NodeProps<NodeData>) {
  const inputRef = useRef<HTMLInputElement>(null);
  // store
  const updateNodeLabel = useStore((state) => state.updateNodeLabel);
  const canDelete = useStore((state) => state.canDelete);
  const deleteNode = useStore((state) => state.deleteNode);
  const confirmDialogState = useConfirmDialogStore();
  // toasts
  const notify = (msg: string) => toast.error(msg);
  const notifySuccess = (msg: string) => toast.success(msg);

  const handleDelete = (id: string) => {
    if (data.root) {
      return;
    }
    confirmDialogState.open(
      "Are you sure you want to delete?",
      () => {
        // Handle confirm action
        try {
          const isDeleteable = canDelete(id);
          if (isDeleteable) {
            // API Request
            deleteNode(id);
            notifySuccess("the node deleted successfully.");
          } else {
            notify("please delete child nodes first.");
          }
        } catch (error) {
          if (error instanceof Error) {
            notify(error.message);
          } else {
            console.log(error);
          }
        }
        // ...perform delete operation
        confirmDialogState.close();
      },
      () => {
        // Handle cancel action
        console.log("Delete canceled");
        confirmDialogState.close();
      }
    );
  };

  useEffect(() => {
    setTimeout(() => {
      inputRef.current?.focus({ preventScroll: true });
    }, 1);
  }, []);

  useLayoutEffect(() => {
    if (inputRef.current) {
      inputRef.current.style.width = `${data.label.length * 8}px`;
    }
  }, [data.label.length]);

  const rootStyle: React.CSSProperties = data.root ? { color: "darkcyan" }: {}
  return (
    <>
      <div className="inputWrapper">
        <div className="dragHandle">
          <DragIcon />
        </div>
        
        <input
          value={data.label}
          style={rootStyle}
          onChange={(evt) => updateNodeLabel(id, evt.target.value)}
          className="input"
          ref={inputRef}
        />
        <div onClick={() => handleDelete(id)}  className="nodeIcon">
          <DeleteIcon Disabled={data.root} />
        </div>
        <div onClick={() => {}} className="nodeIcon">
          <BookIcon />
        </div>
      </div>

      <Handle type="target" position={Position.Top} />
      <Handle type="source" position={Position.Top} />
    </>
  );
}

export default MindMapNode;
