import React from 'react';
import useConfirmDialogStore from './confirmDialogState';

const ConfirmDialog: React.FC = () => {
  const confirmDialogState = useConfirmDialogStore();

  if (!confirmDialogState.isOpen) {
    return null;
  }

  return (
    <div className="dialog">
      <p>{confirmDialogState.message}</p>
      <button className='confirm' onClick={confirmDialogState.onConfirm}>Confirm</button>
      <button className='cancel' onClick={confirmDialogState.onCancel}>Cancel</button>
    </div>
  );
};

export default ConfirmDialog;
