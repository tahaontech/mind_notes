// icon taken from grommet https://icons.grommet.io

type Props = {
    Disabled: boolean
};

const DeleteIcon = ({ Disabled }: Props) => {
  return (
    <svg viewBox="0 0 24 24">
      <path
        fill="none"
        stroke={Disabled ? "#808080" : "#333"}
        strokeWidth="2"
        d="M7.5 9h9v10h-9V9zM5 9h14M9.364 6h5v3h-5V6zm1.181 5v6m3-6v6"
      ></path>
    </svg>
  );
};

export default DeleteIcon;
