import { useDispatch, useSelector } from 'react-redux';

import { actions as slc, IReduxSliceState } from '../ReduxSlice';


// **** ReduxChild **** //

/**
 * Render()
 */
function ReduxChild(): JSX.Element {
    const dispatch = useDispatch();
    const sliceState = useSelector<IReduxSliceState, IReduxSliceState>(sliceState => sliceState);
    return (<div>
       <div>
            Slice Count in Child from Parent Slice: {sliceState.count}
        </div>
        <div>
            <button onClick={() => dispatch(slc.increment())}>
                Incremement
            </button>
        </div>
        <div>
            <button onClick={() => dispatch(slc.decrement())}>
                Decrement
            </button>
        </div>
        <div>
            <button onClick={() => dispatch(slc.reset())}>
                Reset
            </button>
        </div>
    </div>);
}


// Export default.
export default ReduxChild;
