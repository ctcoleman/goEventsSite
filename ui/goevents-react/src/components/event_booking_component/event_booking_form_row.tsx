import * as React from 'react';

export interface FormRowProps {
    label?: string;
}

export class FormRow extends React.Component<FormRowProps, {}> {
    render() {
        return <div className="bx--form-item">
            <label className="bx--col-sm-2 bx--label control-label">
                {this.props.label}
            </label>
            <div className="bx--col-sm-10">
                {this.props.children}
            </div>
        </div>
    }
}