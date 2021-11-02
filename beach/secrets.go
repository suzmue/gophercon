package beach

var sandyspot = beach()

func beach() *hole {
	x := &hole{
		deeper: &hole{
			deeper: &hole{
				deeper: &hole{
					deeper: &hole{
						deeper: &hole{
							deeper: &hole{
								deeper: &hole{
									deeper: &hole{
										deeper: &hole{},
										box:    false,
									},
									box: false,
								},
								box: false,
							},
							box: true,
						},
						box: false,
					},
					box: false,
				},
				box: false,
			},
			box: false,
		},
		box: false,
	}
	return x
}
