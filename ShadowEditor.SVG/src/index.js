import { SvgControl, Manager } from './third_party';

const SVG = new Manager();

window.SVG = SVG;

import './Dom';
import './Circle';
import './Rect';
import './Ellipse';
import './Line';
import './Polyline';
import './Polygon';
import './Path';
import './Text';
import './TextPath';
import './Anchor';

import './defs/Defs';
import './defs/Use';
import './defs/linearGradient';
import './Group';

import './filter/Filter';
import './filter/feGaussianBlur';
import './filter/feOffset';
import './filter/feBlend';
import './filter/feColorMatrix';

export { SvgControl, SVG };