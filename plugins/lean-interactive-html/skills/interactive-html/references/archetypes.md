# Archetype recipes

Per-archetype structure: which skeleton blocks to keep, what to add, and the nontrivial JS recipes verbatim. Always start from `skeleton.html`; these sections tell you how to specialize it.

## report

Audits, PR reviews, stats dashboards, post-mortems.

Keep from skeleton: header + chips, stat cards, badges, bar chart, filter row, expandable table, copy button. Drop: tabs if the report fits one page (use a TOC of in-page anchors instead).

Structure, in order:

1. Header — title, one-line scope, chips with the audit/review meta (date, commit range, reviewer).
2. Stat-card row — the 3–6 numbers that summarize the whole thing.
3. Status table — one row per checked item, `badge ok/warn/bad` verdicts, click-to-expand evidence.
4. Findings as `details.expandable`, severity-badged, collapsed by default, each quoting the actual file/line or command output it came from.
5. Optional "draft response" block with a copy button — e.g. a reply the user can paste to a reviewer or auditor.

## architecture

System or feature architecture documents. The signature element is boxes connected by runtime-computed SVG lines, with a click-to-inspect detail panel.

Keep from skeleton: tokens, header, badges. Add: a `.diagram` container (`position: relative; overflow: hidden`) holding `.row`s of entity boxes and an absolutely-positioned `<svg class="connections">` behind them (`z-index: 1`, boxes at `z-index: 2`, `pointer-events: none` on the svg); a sticky side panel rendering per-entity details from a JS object keyed by `data-key`.

Per-entity accent colors: define `--entityname` and `--entityname-bg` token pairs in `:root` (with dark-mode `-bg` overrides) and a `.block.entityname` class per entity type.

SVG connector recipe — compute endpoints from rendered positions, draw cubic curves with arrowheads and halo-stroked labels, redraw on resize:

```js
const svg = document.getElementById('connections');
const diagram = document.getElementById('diagram');

function getCenter(el) {
  const dr = diagram.getBoundingClientRect();
  const r = el.getBoundingClientRect();
  return {
    cx: r.left + r.width / 2 - dr.left, cy: r.top + r.height / 2 - dr.top,
    top: r.top - dr.top, bottom: r.bottom - dr.top,
    left: r.left - dr.left, right: r.right - dr.left,
  };
}

function drawLines() {
  const dr = diagram.getBoundingClientRect();
  svg.setAttribute('width', dr.width);
  svg.setAttribute('height', dr.height);
  const a = getCenter(document.getElementById('block-a'));
  const b = getCenter(document.getElementById('block-b'));
  const paths = [{
    from: { x: a.cx, y: a.bottom }, to: { x: b.cx, y: b.top },
    cp1: { x: a.cx, y: a.bottom + 60 }, cp2: { x: b.cx, y: b.top - 60 },
    color: '#2563eb', w: 5, label: 'what flows here',
    labelAt: { x: a.cx + 18, y: (a.bottom + b.top) / 2 },
    // dashed: true  — for exception/secondary paths
  }];
  const colors = [...new Set(paths.map(p => p.color))];
  let defs = '<defs>';
  for (const c of colors) {
    defs += `<marker id="arrow-${c.replace('#','')}" viewBox="0 0 10 10" refX="9" refY="5"
      markerWidth="7" markerHeight="7" orient="auto-start-reverse">
      <path d="M 0 0 L 10 5 L 0 10 z" fill="${c}"/></marker>`;
  }
  defs += '</defs>';
  svg.innerHTML = defs + paths.map(p => `
    <path d="M ${p.from.x},${p.from.y} C ${p.cp1.x},${p.cp1.y} ${p.cp2.x},${p.cp2.y} ${p.to.x},${p.to.y}"
          fill="none" stroke="${p.color}" stroke-width="${p.w}" stroke-linecap="round"
          ${p.dashed ? 'stroke-dasharray="12 7"' : ''}
          marker-end="url(#arrow-${p.color.replace('#','')})" />
    <text x="${p.labelAt.x}" y="${p.labelAt.y}" fill="${p.color}" font-size="12.5" font-weight="700"
          paint-order="stroke" stroke="var(--bg)" stroke-width="5">${p.label}</text>`).join('');
}

const scheduleDraw = () => requestAnimationFrame(drawLines);
window.addEventListener('load', scheduleDraw);
window.addEventListener('resize', scheduleDraw);
if (document.fonts && document.fonts.ready) document.fonts.ready.then(scheduleDraw);
```

Route curves around boxes (use a side "lane" `x` for the control points when a straight curve would cross a box). Add a legend in the header mapping line colors to meanings.

Detail panel: a `details` JS object keyed by `data-key`, one `renderPanel(key)` function building `innerHTML` from `{ badge, title, sections: [{ h, l: [...] }] }`, click handler on every `.block` toggling an `.active` class.

## plan

Migration plans, feature plans, briefings — content someone works through over days.

Keep from skeleton: everything. Tabs are mandatory (Overview / Phases / Details / FAQ is a good default), with the keyboard shortcuts the skeleton already wires.

Structure:

1. Overview tab — stat row (scope numbers), goal paragraph, timeline or phase list.
2. One tab or section per phase — each step as a checklist item with a `data-persist` checkbox so progress survives reloads.
3. Filterable + searchable table of work items, click-to-expand rows holding the per-item details.
4. "Evidence from code" collapsibles quoting the real files the plan is based on.

## questionnaire

A set of decisions the user answers in the browser and pastes back into the chat. The artifact IS the feedback loop.

Keep from skeleton: tokens, header, copy helper, localStorage. Drop: tabs, stat cards, table.

Structure: one card per question (number, question, context paragraph, 2–4 clickable options + optional free-text input), and a sticky bottom bar showing live progress ("4/7 answered") plus a Copy button that produces a plain-text summary.

Sticky answer bar recipe:

```js
const answers = JSON.parse(localStorage.getItem(ARTIFACT_KEY + ':answers') || '{}');

function selectOption(qId, value, el) {
  answers[qId] = value;
  localStorage.setItem(ARTIFACT_KEY + ':answers', JSON.stringify(answers));
  el.closest('.options').querySelectorAll('.option').forEach(o => o.classList.remove('selected'));
  el.classList.add('selected');
  renderBar();
}

function renderBar() {
  const total = document.querySelectorAll('.question-card').length;
  const done = Object.keys(answers).length;
  document.getElementById('bar-progress').textContent = `${done}/${total} answered`;
  document.getElementById('answer-summary').innerText =
    Object.entries(answers).map(([q, a]) => `${q}: ${a}`).join('\n');
}
renderBar();
```

The copy button copies `#answer-summary` via the skeleton's `copyText` — plain text, one `question: answer` per line, so the user pastes it straight back into the conversation.

CSS: `.option { cursor: pointer; border: 1px solid var(--border); border-radius: 10px; padding: 10px 14px; }`, `.option.selected { border-color: var(--accent); background: var(--accent-bg); }`, bar is `position: sticky; bottom: 0` with a top border.

## proposal

Process or convention proposals to a team — the reader should be able to *try* the proposed flow, not just read it.

Keep from skeleton: tabs (Overview / Try it / Details / FAQ), collapsibles, table. Add: a current-vs-proposed toggle (two content blocks, one visible at a time) and a step-through simulator.

Ship as a pair: a `.md` file and an `.html` file with the same basename and identical content — the markdown is the reviewable/diffable source, the HTML is the presentation. Commit both together and keep them in sync on every edit.

Step-through simulator recipe — buttons drive a state object; every click appends an explained line to a log:

```js
const START = { /* the initial state of the process being proposed */ };
let S = JSON.parse(JSON.stringify(START));
let step = 1;
const logEl = document.getElementById('log');

function logLine(cls, html) {  // cls: 'info' | 'ok' | 'warnc' | 'err'
  const line = document.createElement('div');
  line.className = 'line';
  line.innerHTML = `<span class="t">${String(step++).padStart(2, '0')}</span><span class="${cls}">${html}</span>`;
  logEl.appendChild(line);
  logEl.scrollTop = logEl.scrollHeight;
}

function render() {
  // reflect S into the state boxes above the log; disable buttons that are
  // invalid in the current state (guards teach the rules better than prose)
}

document.getElementById('action-x').addEventListener('click', () => {
  // mutate S, logLine() each consequence — including the easy-to-forget ones
  render();
});

document.getElementById('action-reset').addEventListener('click', () => {
  S = JSON.parse(JSON.stringify(START));
  step = 1; logEl.innerHTML = '';
  logLine('info', 'Reset. Try: …');
  render();
});

logLine('info', 'Start: … Try clicking the guarded action first — the guard will block it.');
render();
```

Make invalid actions clickable-but-blocked with an explanatory `err` log line — letting the reader trip the guard is the point. Keep a Reset button.
