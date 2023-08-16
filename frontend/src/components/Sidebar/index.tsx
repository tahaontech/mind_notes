import React from 'react';
import './Sidebar.css'; // Import your CSS file

interface SidebarProps {
  children: React.ReactNode;
  onAdd: () => void;
}

const Sidebar: React.FC<SidebarProps> = ({ children, onAdd }) => {
  // const [isSidebarOpen, setIsSidebarOpen] = useState(true);

  // const toggleSidebar = () => {
  //   setIsSidebarOpen(!isSidebarOpen);
  // };

  const isSidebarOpen = true;
  return (
    <div className={`sidebar ${isSidebarOpen ? 'open' : ''}`}>
      <button className="toggle-button" onClick={() => onAdd()}>
        new title
      </button>
      <div className="content">{children}</div>
    </div>
  );
};

export default Sidebar;
