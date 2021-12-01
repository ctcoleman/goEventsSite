import * as React from 'react';
import { Link } from 'react-router-dom';

export interface NavigationProps {
    brandName: string;
}

export class Navigation extends React.Component<NavigationProps, {}> {
    render() {
        return <header className="bx--header" role="banner" aria-label="GoEvents Site" data-header>
            <Link to="/" className="bx--header__name">
                {this.props.brandName}
            </Link>
            <nav className="bx--header__nav" data-header-nav>
                <ul className="bx--header__menu-bar">
                    <li>
                        <Link className="bx--header__menu-item" to="/events">Events</Link>
                    </li>
                </ul>
            </nav>
        </header>
    }
}