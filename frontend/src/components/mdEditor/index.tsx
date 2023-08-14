/* eslint-disable react-hooks/rules-of-hooks */
import React, { useState } from 'react';
import ReactMarkdown from 'react-markdown';
// syntax highlighter
import {Prism as SyntaxHighlighter} from 'react-syntax-highlighter'
import {docco} from 'react-syntax-highlighter/dist/esm/styles/hljs'

import './mdEditor.css'

type Props = {
  id: string;
};

const MDEditore: React.FC<Props> = ({ id }: Props) => {
  const [input, setInput] = useState<string>("")
  console.log(id);
  return (
    <div className="main">
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
