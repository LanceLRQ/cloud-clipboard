import { createBrowserRouter, Navigate } from 'react-router-dom';
import ClipboardBoard from "./clipboard/board";
import ClipboardIndex from "./clipboard";

const RouterIndex = createBrowserRouter([
  {
    path: '/',
    id: 'index',
    element: <Navigate to="/clip" replace={true} />,
  },
  {
    path: '/clip',
    id: 'clip_index',
    element: <ClipboardIndex />,
    children: [
      {
        id: 'clip_board',
        path: 'board',
        element: <ClipboardBoard />
      },
      {
        id: 'clip_gsettins',
        path: 'settings',
        element: <div>Setting</div>
      },
      {
        index: true,
        element: <Navigate to="./board" replace={true} />
      }
    ]
  }
])

export default RouterIndex;