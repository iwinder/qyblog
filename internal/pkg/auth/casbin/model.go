package casbin

const text = `
[request_definition]
r = sub, dom, obj, act

[policy_definition]
p = sub, dom, obj, act

[role_definition]
g = _, _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && g(r.obj, p.obj, r.dom) && (r.dom == p.dom||r.dom=='*') && r.act == p.act
`
