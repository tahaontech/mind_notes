import React from 'react';
import './Sidebar.css'; // Import your CSS file

interface SidebarProps {
  children: React.ReactNode;
}

const Sidebar: React.FC<SidebarProps> = ({ children }) => {
  // const [isSidebarOpen, setIsSidebarOpen] = useState(true);

  // const toggleSidebar = () => {
  //   setIsSidebarOpen(!isSidebarOpen);
  // };

  const isSidebarOpen = true;
  // create category API
  return (
    <div className={`sidebar ${isSidebarOpen ? 'open' : ''}`}>
      <button className="toggle-button" onClick={() => {}}>
        new title
      </button>
      <div className="content">{children}</div>
    </div>
  );
};

export default Sidebar;
