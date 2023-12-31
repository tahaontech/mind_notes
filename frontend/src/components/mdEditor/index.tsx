/* eslint-disable react-hooks/rules-of-hooks */
import React, { useEffect, useState } from 'react';
import ReactMarkdown from 'react-markdown';
// syntax highlighter
import {Prism as SyntaxHighlighter} from 'react-syntax-highlighter'
import {docco} from 'react-syntax-highlighter/dist/esm/styles/hljs'

import './mdEditor.css'
import useEditorStore from './editorState';
import axiosInstance from '../../utils/axiosInstance';
import { toast } from 'react-toastify';

const MDEditore: React.FC = () => {
  // toasts
  const notify = (msg: string) => toast.error(msg);
  const [input, setInput] = useState<string>("");
  const [nodeId, setNodeId] = useState<string>("");
  const editorState = useEditorStore();

  useEffect(() => {
    const id = editorState.message;
    setNodeId(id);
    if (id === "") {
      console.log("fuck react");
    } else {
      (async () => {
        const res = await axiosInstance.get(`/document/${id}`);
        if (res.status === 200) {
          setInput(res.data.data)
        } else {
          notify("thhere is an error while loading");
        }
      })
    }
  }, [editorState.message]);

  const handleSave = async () => {
    const body = { id: nodeId, data: input };
    const res = await axiosInstance.patch("/document", body);
    if (res.status !== 200) {
      notify("there is an error in saving!");
    }
  }

  const handleBack = () => {
    handleSave();
    editorState.close();
  }
  return (
    <div className="main">
      <button onClick={() => handleSave()}>Save</button>
      <button onClick={() => handleBack()}>Back</button>
      <textarea className="textarea" value={input} onChange={(e) => setInput(e.target.value)}/>
      <ReactMarkdown 
        children={input} 
        className='markdown' 
        components={{
          code({ inline, className, children, ...props}) {
            const match = /language-(\w+)/.exec(className || '')
            return !inline && match ? (
              <SyntaxHighlighter
                {...props}
                children={String(children).replace(/\n$/, '')}
                style={docco}
                language={match[1]}
                PreTag="div"
              />
            ) : (
              <code {...props} className={className}>
                {children}
              </code>
            )
          }
        }}
         />
    </div>
  );
};


export default MDEditore;
