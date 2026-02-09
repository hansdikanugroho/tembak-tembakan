components {
  id: "bullet1"
  component: "/main/bullet.script"
}
embedded_components {
  id: "bullet"
  type: "sprite"
  data: "default_animation: \"Bullet-PNG-Clipart\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/sprites.atlas\"\n"
  "}\n"
  ""
  position {
    z: 2.0
  }
  scale {
    x: 0.043151
    y: 0.030701
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_DYNAMIC\n"
  "mass: 1.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"bullet\"\n"
  "mask: \"enemy\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "    id: \"hit\"\n"
  "  }\n"
  "  data: 3.898101\n"
  "  data: 12.815797\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
