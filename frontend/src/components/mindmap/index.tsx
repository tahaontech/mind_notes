import { useCallback, useEffect, useRef } from "react";
import ReactFlow, {
  ConnectionLineType,
  NodeOrigin,
  Node,
  OnConnectEnd,
  OnConnectStart,
  OnMoveEnd,
  useReactFlow,
  useStoreApi,
  Controls,
  Panel,
} from "reactflow";
import { shallow } from "zustand/shallow";

import useStore, { RFState } from "./store";
import MindMapNode from "./MindMapNode";
import MindMapEdge from "./MindMapEdge";

// we need to import the React Flow styles to make it work
import "reactflow/dist/style.css";
import axiosInstance from "../../utils/axiosInstance";
import { toast } from "react-toastify";

type MindmapNodeResp = {
  id: string;
  label: string,
  root: boolean,
  positionX: number,
  positionY: number
}

type MindmapEdgeResp = {
  id: string;
  sourceId: string,
  targetId: string
}

type MindMapResp = {
  category: string;
  nodes: MindmapNodeResp[];
  edges: MindmapEdgeResp[];
}

const selector = (state: RFState) => ({
  nodes: state.nodes,
  edges: state.edges,
  init: state.init,
  onNodesChange: state.onNodesChange,
  onEdgesChange: state.onEdgesChange,
  addChildNode: state.addChildNode,
  deleteNode: state.deleteNode,
});

const nodeTypes = {
  mindmap: MindMapNode,
};

const edgeTypes = {
  mindmap: MindMapEdge,
};

const nodeOrigin: NodeOrigin = [0.5, 0.5];

const connectionLineStyle = { stroke: "#F6AD55", strokeWidth: 3 };
const defaultEdgeOptions = { style: connectionLineStyle, type: "mindmap" };

function Flow({rootId}: {rootId: string;}) {
  const store = useStoreApi();
  const { nodes, edges, onNodesChange, onEdgesChange, addChildNode, init, deleteNode } = useStore(
    selector,
    shallow
  );
  const { project } = useReactFlow();
  const connectingNodeId = useRef<string | null>(null);

  const notify = (msg: string) => toast.error(msg);

  const getChildNodePosition = (event: MouseEvent, parentNode?: Node) => {
    const { domNode } = store.getState();

    if (
      !domNode ||
      // we need to check if these properites exist, because when a node is not initialized yet,
      // it doesn't have a positionAbsolute nor a width or height
      !parentNode?.positionAbsolute ||
      !parentNode?.width ||
      !parentNode?.height
    ) {
      return;
    }

    const { top, left } = domNode.getBoundingClientRect();

    // we need to remove the wrapper bounds, in order to get the correct mouse position
    const panePosition = project({
      x: event.clientX - left,
      y: event.clientY - top,
    });

    // we are calculating with positionAbsolute here because child nodes are positioned relative to their parent
    return {
      x: panePosition.x - parentNode.positionAbsolute.x + parentNode.width / 2,
      y: panePosition.y - parentNode.positionAbsolute.y + parentNode.height / 2,
    };
  };

  const onConnectStart: OnConnectStart = useCallback((_, { nodeId }) => {
    // we need to remember where the connection started so we can add the new node to the correct parent on connect end
    connectingNodeId.current = nodeId;
  }, []);

  const onConnectEnd: OnConnectEnd = useCallback(
    (event) => {
      const { nodeInternals } = store.getState();
      const targetIsPane = (event.target as Element).classList.contains(
        "react-flow__pane"
      );
      const node = (event.target as Element).closest(".react-flow__node");

      if (node) {
        node.querySelector("input")?.focus({ preventScroll: true });
      } else if (
        targetIsPane &&
        connectingNodeId.current &&
        event instanceof MouseEvent
      ) {
        const parentNode = nodeInternals.get(connectingNodeId.current);
        const childNodePosition = getChildNodePosition(event, parentNode);

        if (parentNode && childNodePosition) {
          const data = addChildNode(parentNode, childNodePosition);
          // API: call craete node 
          const callBody = {
            id: data.node.id,
            label: data.node.data.label,
            root: false,
            positionX: data.node.position.x,
            positionY: data.node.position.y,
            rootId: data.node.data.rootId,
            edgeId: data.edge.id,
            sourceId: data.edge.source
          };
          (async () => {
            try {
              const res = await axiosInstance.post("/node", callBody);
              if (res.status !== 200) {
                notify("node not added.");
                deleteNode(data.node.id);
              }
            } catch (error) {
              console.log(error);
              notify("node not added.");
              deleteNode(data.node.id);
            }
          })();
          const node = (event.target as Element).closest(".react-flow__node");
          if (node) {
            node.querySelector("input")?.focus({ preventScroll: true });
          }
        }
      }
    },
    [getChildNodePosition]
  );

  // TODO: add moveEnd event to update position
  const onPositionChanged: OnMoveEnd = useCallback(
    (event: MouseEvent | TouchEvent) => {
      if (event instanceof MouseEvent) {
        console.log("node movee", event.clientX)
      }
    },
    []
  )

  useEffect(() => {
    // retrive nodes and edges based on rootId
    if (rootId === "") return;
    (async () => {
      try {
        const resp = await axiosInstance.get<MindMapResp>(`/mindmap/${rootId}`)
        if (resp.status === 200) {
          const nodes = resp.data.nodes.map((v) => { 
            return {
              id: v.id,
              type: "mindmap", // static in client
              data: { label: v.label, root: v.root, rootId: rootId },
              position: { x: v.positionX, y: v.positionY },
              dragHandle: ".dragHandle",
            }
          })
          const edges = resp.data.edges.map((v) => {
            return {
              id: v.id,
              source: v.sourceId,
              target: v.targetId,
            }
          })
          init(nodes, edges)
        } else {
          notify(`server error: ${resp.status}`)
        }
      } catch (error) {
        notify("there is an error in mindmap")
      }
    })()
  }, [rootId])
  

  return (
    <ReactFlow
      nodes={nodes}
      edges={edges}
      onNodesChange={onNodesChange}
      onEdgesChange={onEdgesChange}
      onConnectStart={onConnectStart}
      onConnectEnd={onConnectEnd}
      onMoveEnd={onPositionChanged}
      nodeTypes={nodeTypes}
      edgeTypes={edgeTypes}
      nodeOrigin={nodeOrigin}
      defaultEdgeOptions={defaultEdgeOptions}
      connectionLineStyle={connectionLineStyle}
      connectionLineType={ConnectionLineType.Straight}
      fitView
    >
      <Controls showInteractive={false} />
      <Panel position="top-left" className="header">
        Mind Map
      </Panel>
    </ReactFlow>
  );
}

export default Flow;
