import {
  Edge,
  EdgeChange,
  Node,
  NodeChange,
  OnNodesChange,
  OnEdgesChange,
  applyNodeChanges,
  applyEdgeChanges,
  XYPosition,
} from "reactflow";
import create from "zustand";
import { nanoid } from "nanoid/non-secure";

import { NodeData } from "./MindMapNode";

export type RFState = {
  nodes: Node<NodeData>[];
  edges: Edge[];
  onNodesChange: OnNodesChange;
  onEdgesChange: OnEdgesChange;
  updateNodeLabel: (nodeId: string, label: string) => void;
  deleteNode: (nodeId: string) => void;
  addChildNode: (parentNode: Node, position: XYPosition) => void;
  canDelete: (nodeId: string) => boolean;
};

const useStore = create<RFState>((set, get) => ({
  nodes: [
    {
      id: "root",
      type: "mindmap",
      data: { label: "root node", root: true },
      position: { x: 0, y: 0 },
      dragHandle: ".dragHandle",
    },
  ],
  edges: [],
  onNodesChange: (changes: NodeChange[]) => {
    set({
      nodes: applyNodeChanges(changes, get().nodes),
    });
  },
  onEdgesChange: (changes: EdgeChange[]) => {
    set({
      edges: applyEdgeChanges(changes, get().edges),
    });
  },
  updateNodeLabel: (nodeId: string, label: string) => {
    set({
      nodes: get().nodes.map((node) => {
        if (node.id === nodeId) {
          // it's important to create a new object here, to inform React Flow about the changes
          node.data = { ...node.data, label };
        }

        return node;
      }),
    });
  },
  canDelete : (nodeId: string) => {
    const node = get().nodes.filter(n => n.id === nodeId)[0];
    const childEdges = get().edges.filter(e => e.source === nodeId);
    if (node.data.root === false && childEdges.length === 0) {
      return true;
    }
    return false;
  },
  deleteNode: (nodeId: string) => {
    const node = get().nodes.filter(n => n.id === nodeId)[0];
    const childEdges = get().edges.filter(e => e.source === nodeId);
    if (node.data.root === false && childEdges.length === 0) {
      set({
        nodes: get().nodes.filter(n => n.id !== nodeId ),
        edges: get().edges.filter(e => e.target !== nodeId)
      })
    } else {
      throw new Error('please delete child nodes first.')
    }
  },
  addChildNode: (parentNode: Node, position: XYPosition) => {
    const newNode = {
      id: nanoid(),
      type: "mindmap",
      data: { label: "New Node", root: false },
      position,
      dragHandle: ".dragHandle",
      parentNode: parentNode.id,
    };

    const newEdge = {
      id: nanoid(),
      source: parentNode.id,
      target: newNode.id,
    };

    set({
      nodes: [...get().nodes, newNode],
      edges: [...get().edges, newEdge],
    });
  },
}));

export default useStore;
