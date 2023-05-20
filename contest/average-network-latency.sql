SELECT CAST(
    (
        (SELECT sum(
                    1000 * EXTRACT (EPOCH FROM rr.datetime) - 1000 * EXTRACT (EPOCH FROM rs.datetime)
                )   /   (SELECT count(DISTINCT request_id)
                        FROM requests
                        WHERE parent_request_id IS NULL)
                    FROM requests rs
                    INNER JOIN requests rr ON rs.request_id=rr.parent_request_id
                    AND rr.host=rs.data
                    AND rs.type='RequestSent'
                    AND rr.type='RequestReceived'
        ) +
            (SELECT sum(1000 * EXTRACT (EPOCH FROM rr.datetime) - 1000 * EXTRACT (EPOCH FROM rs.datetime))/
            (SELECT count(DISTINCT request_id)
                 FROM requests
                 WHERE parent_request_id IS NULL
             )
             FROM requests rs
             INNER JOIN requests rr ON rr.request_id=rs.parent_request_id
             AND rr.data like concat(rs.host, '%')
             AND rs.type='ResponseSent'
             AND rr.type='ResponseReceived'
        )
    ) AS NUMERIC
) AS avg_network_time_ms